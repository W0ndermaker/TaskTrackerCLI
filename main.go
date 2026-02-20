package main

import (
	"fmt"
	"os"
	"strings"

	"TaskTrackerCLI/commands"
	"TaskTrackerCLI/internal"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Invalid command: Not enough arguments")
		os.Exit(1)
	}

	// reading storage file
	tasksFromStorage, err := internal.ReadDataFromStorage(commands.StorageName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// lowercasing command and args
	command, restArgs := strings.ToLower(os.Args[1]), os.Args[2:]

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
		if len(restArgs) != 1 {
			fmt.Println("Invalid command: incorrect amount of args for -delete- command")
			os.Exit(1)
		}
		err := commands.Delete(restArgs[0], tasksFromStorage)
		if err != nil {
			fmt.Println(err)
		}
	case "mark-in-progress":
		if len(restArgs) != 1 {
			fmt.Println("Invalid command: incorrect amount of args for -mark_in_progress- command")
			os.Exit(1)
		}
		err := commands.Mark(restArgs[0], "in-progress", tasksFromStorage)
		if err != nil {
			fmt.Println(err)
		}
	case "mark-done":
		if len(restArgs) != 1 {
			fmt.Println("Invalid command: incorrect amount of args for -mark_done- command")
			os.Exit(1)
		}
		err := commands.Mark(restArgs[0], "done", tasksFromStorage)
		if err != nil {
			fmt.Println(err)
		}
	case "list":
		statusFilter := ""
		if len(restArgs) == 1 {
			statusFilter = restArgs[0]
		}

		if len(restArgs) > 1 {
			fmt.Println("Invalid command: incorrect amount of args for -list- command")
			os.Exit(1)
		}
		err := commands.List(statusFilter, tasksFromStorage)
		if err != nil {
			fmt.Println(err)
		}

	default:
		fmt.Println("Unknown command")
	}
}
