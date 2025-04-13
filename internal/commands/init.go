// internal/commands/init.go
package commands

import (
    "fmt"
    "os"
    "path/filepath"
)

func InitRepo() {
    
    gitlDir := ".gitl"
    // Check if already initialized
    stat, err := os.Stat(gitlDir)
    if err == nil {
        fmt.Println("[!] gitl exists already!...")
        return
    } else if !os.IsNotExist(err) {
        fmt.Println(">> Error accessing .gitl:", err)
        return
    } else {
        fmt.Println("[✓] .gitl is perfectly setup!... ")
    }

    // Create .gitl structure
    dirs := []string{
        filepath.Join(gitlDir, "objects"),
        filepath.Join(gitlDir, "refs", "heads"),
        filepath.Join(gitlDir, "refs", "tags"),
    }

    for _, dir := range dirs {
        fmt.Println(">> Creating dir:", dir)
        dirErr := os.MkdirAll(dir, 0755) // Use dirErr to avoid shadowing
        if dirErr != nil {
            //fmt.Printf(">> Error creating directory %s: %v\n", dir, dirErr)
            return
        }
    }

    headPath := filepath.Join(gitlDir, "HEAD")
    fileErr := os.WriteFile(headPath, []byte("ref: refs/heads/master\n"), 0644) // Use fileErr
    if fileErr != nil {
        //fmt.Printf(">> Error writing HEAD: %v\n", fileErr)
        return
    }

    //fmt.Println(">> Successfully initialized Gitl repository in .gitl/")
}
