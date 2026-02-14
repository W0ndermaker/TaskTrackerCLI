package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func Add(TaskName string) error {
	// открытие файла для чтения и записи если существует, если нет то создаётся файл
	file, err := os.OpenFile("./storage.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error occured while creating storage file:")
		return err
	}
	defer file.Close()

	// надо прочитать, десериализировать и взять номер Id

	//...

	// создание таски
	t := task{
		ID:          0,
		Description: TaskName,
		CreatedAt:   "",
		UpdatedAt:   "",
	}
	// сериализация
	data, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		fmt.Println("Error occured while marshalising the data into JSON")
		return err
	}

	// чтение самого первого элемента файла
	buf := make([]byte, 1)
	n, err := file.Read(buf)

	// запись
	if n == 0 {
		file.WriteString("[" + string(data) + "]")
	} else {
		file.Seek(-1, io.SeekEnd)
		file.WriteString(",\n" + string(data) + "]")
	}

	fmt.Println("Task was added")

	return nil
}

func Update() {}

func Delete() {}

func Mark_in_progress() {}

func Mark_done() {}

func List() {}
