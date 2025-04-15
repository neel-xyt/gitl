package main

import (
	//"fmt"
	"gitl/internal/VersionControl" // Import the VersionControl package
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: gitl <command>")
		return
	}

	switch os.Args[1] {
	case "init":
		versioncontrol.Init()	
		versioncontrol.Gmanager()
	default:
		println("Unknown command:", os.Args[1])
	}
}
