package commands

import (
	"fmt"
	"strings"
	"time"

	"TaskTrackerCLI/internal"
)

func Add(TaskName string, tasks []internal.Task) error {
	// Поиск Id последней добавленной таски
	lastTaskId := 1
	if len(tasks) != 0 {
		lastTaskId += tasks[len(tasks)-1].ID
	}
	// Время добавления таски
	creadtedTime := time.Now()
	formattedTime := creadtedTime.Format("2006-01-02 15:04:05")

	// создание таски
	newTask := internal.Task{
		ID:          lastTaskId,
		Description: TaskName,
		Status:      "todo",
		CreatedAt:   formattedTime,
		UpdatedAt:   formattedTime,
	}

	tasks = append(tasks, newTask)

	err := internal.WriteToStorage(tasks)
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

	// Сохраняем старое значение для логирования
	oldDescription := tasks[taskIndex].Description

	//  Обновление задачи
	tasks[taskIndex].Description = newTaskName

	// Сохранение в хранилище
	if err = internal.WriteToStorage(tasks); err != nil {
		// Восстанавливаем предыдущее значение в случае ошибки
		tasks[taskIndex].Description = oldDescription
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

	err = internal.WriteToStorage(tasks)
	if err != nil {
		return fmt.Errorf("failed to save updated task: %w", err)
	}
	fmt.Printf("Task(ID:%v) was deleted successfully\n", id)
	return nil
}

func Mark(id string, tasks []internal.Task, newStatus string) error {
	taskIndex, err := internal.FindTaskID(id, tasks)
	if err != nil {
		return err
	}

	// Если статус задачи уже такой же как новый статус
	if tasks[taskIndex].Status == newStatus {
		fmt.Printf("Status of this task(id:%v) is already %v", id, newStatus)
		return nil
	}
	// Сохраняем старое значение для логирования
	oldStatus := tasks[taskIndex].Status

	//  Обновление задачи
	tasks[taskIndex].Status = newStatus

	// Сохранение в хранилище
	if err = internal.WriteToStorage(tasks); err != nil {
		// Восстанавливаем предыдущее значение в случае ошибки
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
