package vcs

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
		"branches.json",
		"COMMIT_EDITMSG.json",
		"config.json",
		"description.json",
		"HEAD.json",
		"hooks.json",
		"index.json",
		"info/exclude.json",
		"logs/HEAD.json",
		"objects.json",
		"refs/heads.json",
		"refs/tags.json",
		"refs/remotes.json",
		"packed-refs.json",
	}

	// Create the main .gitl directory
	if err := os.MkdirAll(gitlDir, os.ModePerm); err != nil {
		return
	}

	for _, file := range files {
		filePath := filepath.Join(gitlDir, file)

		// Create parent directories if needed
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			continue
		}

		// Write empty JSON object to file
		if err := os.WriteFile(filePath, []byte("{}"), 0644); err != nil {
			continue
		}
	}
}
