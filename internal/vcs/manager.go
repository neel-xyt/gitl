package vcs

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

// AES encryption function using CFB128 mode
func encrypt(data []byte, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("Oops! Something went wrong while encrypting the data: %v", err)
	}

	// Prepare the ciphertext slice with space for IV and data
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", fmt.Errorf("Failed to generate random bytes for encryption: %v", err)
	}

	// Use the CFB mode (CFB128) with the initialization vector
	stream := cipher.NewCFB(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	// Return base64-encoded ciphertext for storage
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// GenerateEncryptedJSON function to save encrypted JSON to file
func GenerateEncryptedJSON(data any, filePath, key string) {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0700); err != nil {
		log.Printf("Failed to create directory %s: %v\n", dir, err)
		log.Println("Please ensure you have the required permissions.")
		return
	}

	// Marshal the data into JSON format
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("Error converting data to JSON: %v\n", err)
		log.Println("Please check the structure of your data.")
		return
	}

	// Encrypt the JSON data
	encryptedData, err := encrypt(jsonData, key)
	if err != nil {
		log.Printf("Error encrypting data: %v\n", err)
		log.Println("Please try again later.")
		return
	}

	// Write the encrypted data to the file with restricted permissions (0600)
	if err := os.WriteFile(filePath, []byte(encryptedData), 0600); err != nil {
		log.Printf("Unable to save data to %s: %v\n", filePath, err)
		log.Println("Make sure the directory exists and you have proper write permissions.")
		return
	}

	fmt.Println("[+] User configuration saved successfully (encrypted).")
}

type Parson struct {
	Bare                     bool    `json:"bare"`
	RepositoryFormatVersion float32 `json:"repositoryFormatVersion"`
	FileMode                 bool    `json:"fileMode"`
}

type ConfigWrapped struct {
	Coreinfo Parson `json:"core"`
}

// Main function for generating configuration
func Gmanager() {
	var P Parson
	P.Bare = false
	P.RepositoryFormatVersion = 0.1
	P.FileMode = runtime.GOOS != "windows"
	corei := ConfigWrapped{Coreinfo: P}

	// Get the home directory to store the configuration file in a hidden .gitl directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Couldn't retrieve the home directory: %v\n", err)
		log.Println("Please ensure your system is configured properly.")
		return
	}

	gitlDataDir := filepath.Join(homeDir, ".gitl", "data")
	os.MkdirAll(gitlDataDir, 0700)

	filePath := filepath.Join(gitlDataDir, "config.json")

	// Save the configuration file securely
	GenerateEncryptedJSON(corei, filePath, "your-encryption-key")
}

type UserInfo struct {
	Username string `json:"username"`
	Useremil string `json:"useremil"`
}

type User struct {
	User UserInfo `json:"userdata"`
}

func isValidFormat(useremil string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(useremil)
}

func domainHasMX(useremil string) bool {
	parts := strings.Split(useremil, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	mx, err := net.LookupMX(domain)
	return err == nil && len(mx) > 0
}

func isRealEmail(useremil string) bool {
	return isValidFormat(useremil) && domainHasMX(useremil)
}

func Firstsetup() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("[?] Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("[?] Enter your email: ")
	useremil, _ := reader.ReadString('\n')
	useremil = strings.TrimSpace(useremil)

	if isRealEmail(useremil) {
		finalData := User{
			User: UserInfo{
				Username: username,
				Useremil: useremil,
			},
		}

		// Encrypt and save the user configuration
		homeDir, _ := os.UserHomeDir()
		gitlDataDir := filepath.Join(homeDir, ".gitl", "data")
		os.MkdirAll(gitlDataDir, 0700)

		filePath := filepath.Join(gitlDataDir, "userinfo.json")

		// Save the configuration file securely
		GenerateEncryptedJSON(finalData, filePath, "your-encryption-key")
		fmt.Println("[+] User configuration saved successfully (encrypted).")
	} else {
		fmt.Printf("[-] The email %s is invalid or domain doesn't exist.\n", useremil)
	}
}
