package main

import (
	"fmt"
	"os"

	"TaskTrackerCLI/commands"
	"TaskTrackerCLI/internal"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Invalid command: Not enough arguments")
		os.Exit(1)
	}

	tasksFromStorage, err := internal.ReadDataFromStorage()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	command, restArgs := os.Args[1], os.Args[2:]

	switch command {
	case "add":
		if len(restArgs) != 1 {
			fmt.Println("Invalid command: incorrect amount of args for -add- command")
			os.Exit(1)
		}
		err := commands.Add(restArgs[0], tasksFromStorage)
		if err != nil {
			fmt.Println(err)
		}
	case "update":
		if len(restArgs) != 2 {
			fmt.Println("Invalid command: incorrect amount of args for -update- command")
			os.Exit(1)
		}
		err := commands.Update(restArgs[0], restArgs[1], tasksFromStorage)
		if err != nil {
			fmt.Println(err)
		}
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
