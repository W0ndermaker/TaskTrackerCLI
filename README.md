# TaskTrackerCLI
==============

**TaskTrackerCLI** is a simple command-line task tracker written in Go.  
Tasks are stored in a local JSON file (`storage.json`), and you can add, update, delete, and list them directly from your terminal.

### Installation

- **Prerequisites**: Go installed and available on your `PATH`.

From the project root:

```bash
go build -o TaskTrackerCLI
```

This will create an executable named `TaskTrackerCLI` in the project directory.

### Usage

Run the CLI by calling the executable followed by a command and its arguments:

```bash
TaskTrackerCLI <command> [arguments]
```

If you run it without any arguments, it will print an error about missing arguments.

### Commands

- **add**

  Add a new task with a description.

  ```bash
  TaskTrackerCLI add "Buy groceries"
  ```

- **update**

  Update the description of an existing task.

  ```bash
  TaskTrackerCLI update <taskID> "New description"
  ```

- **delete**

  Delete a task by its ID.

  ```bash
  TaskTrackerCLI delete <taskID>
  ```

- **mark-in-progress**

  Mark a task as in progress.

  ```bash
  TaskTrackerCLI mark-in-progress <taskID>
  ```

- **mark-done**

  Mark a task as done.

  ```bash
  TaskTrackerCLI mark-done <taskID>
  ```

- **list**

  List tasks, optionally filtered by status.

  ```bash
  # List all tasks
  TaskTrackerCLI list

  # List only tasks with a specific status:
  # allowed values depend on your internal implementation (e.g. "todo", "in-progress", "done")
  TaskTrackerCLI list done
  ```

### Data Storage

- Tasks are loaded from and saved to `storage.json` in the project directory.
- If there are problems reading this file, the program will print an error and exit.

### Exit Codes

- **0**: Success (command ran without fatal errors).
- **1**: Invalid command usage or storage error.
