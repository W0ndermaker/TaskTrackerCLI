package commands

import (
	"fmt"
	"strconv"
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

	if !internal.IsValidID(id) {
		return fmt.Errorf("invalid taskID format: %s", id)
	}

	convertedID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("failed to convert taskID '%s' to integer: %w", id, err)
	}

	//  Поиск задачи (бинарынй поиск)
	taskIndex := internal.BinarySearchOfTaskID(tasks, convertedID)
	if taskIndex == -1 {
		return fmt.Errorf("task with ID %d not found", convertedID)
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

func Delete() {}

func Mark_in_progress() {}

func Mark_done() {}

func List() {}
