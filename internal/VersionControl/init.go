// internal/commands/init.go
package versioncontrol

import (
    "fmt"
    "os"
    "path/filepath"
)
 func Init() {
    gitDir := ".gitl"
    _, err := os.Stat(gitDir)
    if err != nil { // if the .gitl directory does not exist
        GitldirC()
        fmt.Println("[ok] Gitl setup...")
    } else {
        fmt.Println("[!] The .gitl directory already exists...")
    }
}


func GitldirC() {
    gitlDir := ".gitl"

    files := []string{
        "branches.yml",
        "COMMIT_EDITMSG.yml",
        "config.yml",
        "description.yml",
        "HEAD.yml",
        "hooks.yml",
        "index.yml",
        "info/exclude.yml",
        "logs/HEAD.yml",
        "objects.yml",
        "refs/heads.yml",
        "refs/tags.yml",
        "refs/remotes.yml",
        "packed-refs.yml",
    }

    // Create the directory if it doesn't exist
    if err := os.MkdirAll(gitlDir, os.ModePerm); err != nil {
        return
    }

    for _, file := range files {
        filePath := filepath.Join(gitlDir, file)

        // Create parent directories for the file if they don't exist
        if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
            continue
        }

        f, err := os.Create(filePath)
        if err != nil {
            continue
        }
        defer f.Close()
    }
}
