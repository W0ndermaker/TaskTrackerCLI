package commands

import (
	"fmt"
	"strings"
	"time"

	"TaskTrackerCLI/internal"
)

var StorageName = "./storage.json"

func Add(taskName string, tasks []internal.Task) error {
	lastTaskId := 1
	if len(tasks) != 0 {
		lastTaskId += tasks[len(tasks)-1].ID
	}
	creadtedTime := time.Now().Format("2006-01-02 15:04:05")

	newTask := internal.Task{
		ID:          lastTaskId,
		Description: taskName,
		Status:      "todo",
		CreatedAt:   creadtedTime,
		UpdatedAt:   creadtedTime,
	}

	tasks = append(tasks, newTask)

	err := internal.WriteToStorage(StorageName, tasks)
	if err != nil {
		return fmt.Errorf("failed to save updated task: %w", err)
	}
	fmt.Println("New task was added successfully!")
	return nil
}

func Update(id string, newTaskName string, tasks []internal.Task) error {
	if strings.TrimSpace(newTaskName) == "" {
		return fmt.Errorf("task name cannot be empty")
	}

	taskIndex, err := internal.FindTaskID(id, tasks)
	if err != nil {
		return err
	}

	oldDescription := tasks[taskIndex].Description
	oldUpdatedTime := tasks[taskIndex].UpdatedAt

	tasks[taskIndex].Description = newTaskName
	tasks[taskIndex].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	if err = internal.WriteToStorage(StorageName, tasks); err != nil {
		tasks[taskIndex].Description = oldDescription
		tasks[taskIndex].UpdatedAt = oldUpdatedTime
		return fmt.Errorf("failed to save updated task: %w", err)
	}
	fmt.Printf("Task(ID:%v) was updated successfully\n", tasks[taskIndex].ID)
	return nil
}

func Delete(id string, tasks []internal.Task) error {
	taskIndex, err := internal.FindTaskID(id, tasks)
	if err != nil {
		return err
	}
	tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)

	err = internal.WriteToStorage(StorageName, tasks)
	if err != nil {
		return fmt.Errorf("failed to save updated task: %w", err)
	}
	fmt.Printf("Task(ID:%v) was deleted successfully\n", id)
	return nil
}

func Mark(id string, newStatus string, tasks []internal.Task) error {
	taskIndex, err := internal.FindTaskID(id, tasks)
	if err != nil {
		return err
	}

	if tasks[taskIndex].Status == newStatus {
		fmt.Printf("Status of this task(id:%v) is already %v", id, newStatus)
		return nil
	}

	oldStatus := tasks[taskIndex].Status

	tasks[taskIndex].Status = newStatus

	if err = internal.WriteToStorage(StorageName, tasks); err != nil {
		tasks[taskIndex].Status = oldStatus
		return fmt.Errorf("failed to save updated task: %w", err)
	}
	fmt.Printf("Status of task(ID:%v) was updated successfully\n", tasks[taskIndex].ID)
	return nil
}

func List(filter string, tasks []internal.Task) error {
	filter = strings.TrimSpace(filter)
	if filter == "" {
		for _, task := range tasks {
			internal.PrintTask(task)
		}
		return nil
	}
	if filter != "todo" && filter != "in-progress" && filter != "done" {
		return fmt.Errorf("Unknown status")
	}

	foundTask := false
	for _, task := range tasks {
		if task.Status == filter {
			foundTask = true
			internal.PrintTask(task)
		}
	}

	if !foundTask {
		fmt.Printf("No tasks with status:%v\n", filter)
	}

	return nil
}
