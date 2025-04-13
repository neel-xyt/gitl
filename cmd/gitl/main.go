package main

import (
    "os"
    "gitl/internal/commands" // Import the commands package
)

func main() {
    if len(os.Args) < 2 {
        println("Usage: gitl <command>")
        return
    }

    switch os.Args[1] {
    case "init":
        commands.InitRepo() // Call the InitRepo function from the commands package
    default:
        println("Unknown command:", os.Args[1])
    }
}
