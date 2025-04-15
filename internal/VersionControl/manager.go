package versioncontrol

import (
	"encoding/json"
	//"fmt"
	"log"
	"os"
	"runtime"
	"path/filepath"
)

func GenerateJSON(data any, filePath string) {
    // Ensure the directory exists
    dir := filepath.Dir(filePath)
    if err := os.MkdirAll(dir, os.ModePerm); err != nil {
        log.Fatal("Error creating directory:", err)
    }

    // Marshal the data into JSON format
    jsonData, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        log.Fatal("Error marshaling JSON:", err)
    }

    // Write the JSON data to the specified file
    err = os.WriteFile(filePath, jsonData, 0644)
    if err != nil {
        log.Fatal("Error writing to file:", err)
    }
}
type Parson struct {
    Bare                    bool    `json:"bare"`
    RepositoryFormatVersion float32 `json:"repositoryFormatVersion"`
    FileMode                bool    `json:"fileMode"`
}
type ConfigWrapped struct {
	Coreinfo Parson `json:"core"`
}
func Gmanager() {
 		var P Parson 
	P.Bare = false
	P.RepositoryFormatVersion = 0.1
	P.FileMode = runtime.GOOS != "windows"
	corei := ConfigWrapped {Coreinfo: P}
	gitlpath := ".gitl/config.json"
	GenerateJSON(corei,gitlpath)
} 
