package vcs

import (
	"bufio"
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

func GenerateJSON(data any, filePath string) {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal("Error creating directory:", err)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}

	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		log.Fatal("Error writing to file:", err)
	}
}

type Parson struct {
	Bare                     bool    `json:"bare"`
	RepositoryFormatVersion float32 `json:"repositoryFormatVersion"`
	FileMode                 bool    `json:"fileMode"`
}

type ConfigWrapped struct {
	Coreinfo Parson `json:"core"`
}

func Gmanager() {
	var P Parson
	P.Bare = false
	P.RepositoryFormatVersion = 0.1
	P.FileMode = runtime.GOOS != "windows"
	corei := ConfigWrapped{Coreinfo: P}
	gitlpath := ".gitl/config.json"
	GenerateJSON(corei, gitlpath)
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
		GenerateJSON(finalData, "gitl/data/userinfo.json")
		fmt.Println("[+] User configuration saved successfully.")
	} else {
		fmt.Printf("[-] The email %s is invalid or domain doesn't exist.\n", useremil)
	}
}
