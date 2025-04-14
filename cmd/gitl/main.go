package main

import (
	"os"
	"gitl/internal/VersionControl" // Import the VersionControl package
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: gitl <command>")
		return
	}

	switch os.Args[1] {
	case "init":
		versioncontrol.Init() // Call the InitRepo function from the VersionControl package
	default:
		println("Unknown command:", os.Args[1])
	}
}
