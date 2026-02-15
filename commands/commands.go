package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"TaskTrackerCLI/commands/inner"
)

type task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

const storageName = "./storage.json"

func Add(TaskName string) error {
	// открытие файла для чтения и записи если существует, если нет то создаётся файл
	file, err := os.OpenFile(storageName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error occured while creating storage file:")
		return err
	}
	defer file.Close()

	// чтение данных уже находящихся в файле
	dataFromFile, err := inner.ReadData(storageName)
	if err != nil {
		return err
	}
	fileDataLen := len(dataFromFile)

	// Поиск Id последней добавленной таски
	lastTaskId := 0
	tasks := []task{}
	if fileDataLen != 0 {

		err = json.Unmarshal(dataFromFile, &tasks)
		if err != nil {
			fmt.Println("Error occurred wgile unmarshalling")
			return err
		}

		lastTaskId = tasks[len(tasks)-1].ID
	}

	for _, t := range tasks {
		if t.Description == TaskName {
			fmt.Println("This task already exists!")
			return nil
		}
	}

	// Время добавления таски
	creadtedTime := time.Now()
	formattedTime := creadtedTime.Format("2006-01-02 15:04:05")

	// создание таски
	t := task{
		ID:          lastTaskId + 1,
		Description: TaskName,
		Status:      "todo",
		CreatedAt:   formattedTime,
		UpdatedAt:   formattedTime,
	}
	// сериализация
	data, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		fmt.Println("Error occured while marshalising the data into JSON")
		return err
	}

	// запись
	if fileDataLen == 0 {
		file.WriteString("[" + string(data) + "]")
	} else {
		file.Seek(-1, io.SeekEnd)
		file.WriteString(",\n" + string(data) + "]")
	}

	fmt.Println("Task was added successfully!!!")

	return nil
}

func Update() {}

func Delete() {}

func Mark_in_progress() {}

func Mark_done() {}

func List() {}
