package commands

import (
	"encoding/json"
	"fmt"
	"os"
)

type task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func newTask(id int, description string, status string, createdAt string, updatedAt string) *task {
	return &task{id, description, status, createdAt, updatedAt}
}

func Add(TaskName string) error {
	file, err := os.OpenFile("./storage.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error occured while creating storage file:")
		return err
	}
	defer file.Close()

	t := newTask(0, TaskName, "", "", "")

	data, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Error occured while marshalising the data into JSON")
		return err
	}

	file.WriteString(string(data) + "\n")
	fmt.Println("Task was added")

	return nil
}

func Update() {}

func Delete() {}

func Mark_in_progress() {}

func Mark_done() {}

func List() {}
