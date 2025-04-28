package vcs

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"io/fs"
)

type Change struct {
	Path string
	Type string // "new", "modified", "directory"
}

type TrackedFile struct {
	Hash string
}

func hashFile(path string) (string, error) {
	data, err := os.ReadFile(path) // Using os.ReadFile (Go 1.16+)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil
}

func TrackChanges(root string, tracked map[string]TrackedFile) ([]Change, error) {
	var changes []Change

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || filepath.Base(path) == ".gitl" {
			return nil
		}

		relPath, _ := filepath.Rel(root, path)

		if d.IsDir() {
			if relPath != "." && tracked[relPath].Hash == "" {
				changes = append(changes, Change{Path: relPath, Type: "directory"})
			}
			return nil
		}

		newHash, err := hashFile(path)
		if err != nil {
			return nil
		}

		old, exists := tracked[relPath]
		if !exists {
			changes = append(changes, Change{Path: relPath, Type: "new"})
		} else if old.Hash != newHash {
			changes = append(changes, Change{Path: relPath, Type: "modified"})
		}

		return nil
	})

	return changes, err
}
