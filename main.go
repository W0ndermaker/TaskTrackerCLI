package main

import (
	"fmt"
	"os"

	"TaskTrackerCLI/commands"
)

func main() {
	if len(os.Args[1:]) < 2 {
		fmt.Println("Invalid command: Not enough arguments")
		os.Exit(1)
	}

	command := os.Args[1]
	cmdArg1 := os.Args[2]
	switch command {
	case "add":
		commands.Add(cmdArg1)
	case "update":
		commands.Update()
	case "delete":
		commands.Delete()
	case "mark-in-progress":
		commands.Mark_in_progress()
	case "mark-done":
		commands.Mark_done()
	case "list":
		commands.List()
	default:
		fmt.Println("Unknown command")
	}
}
