package main

import (
	//"fmt"
	"gitl/internal/vcs"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: gitl <command>")
		return
	}

	switch os.Args[1] {
	case "init":
		vcs.Init()	
		vcs.Gmanager()
	case "fsetup": 
		vcs.Firstsetup()
	default:
		println("Unknown command:", os.Args[1])
	}
}
