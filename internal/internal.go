package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const storageName = "./storage.json"

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func ReadDataFromStorage() ([]Task, error) {
	// открытие файла для чтения если существует, если нет то создаётся файл
	file, err := os.OpenFile(storageName, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error occured while creating storage file:")
		return nil, err
	}
	defer file.Close()

	// подсчёт размера файла для буфера
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()

	// если файл изначально пустой
	if fileSize == 0 {
		return []Task{}, nil
	}

	myBuf := make([]byte, fileSize)

	// чтение данных уже находящихся в файле
	_, err = file.Read(myBuf)
	if err != nil {
		fmt.Println("Error while reading file")
		return nil, err
	}

	tasks := []Task{}
	err = json.Unmarshal(myBuf, &tasks)
	if err != nil {
		fmt.Println("Error occurred wgile unmarshalling:")
		return nil, err
	}
	return tasks, nil
}

func WriteToStorage(newInfo []Task) error {
	// открытие файла для записи
	file, err := os.OpenFile(storageName, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("Error occured while creating storage file: %w", err)
	}
	defer file.Close()
	// сериализация
	data, err := json.MarshalIndent(newInfo, "", " ")
	if err != nil {
		return fmt.Errorf("Error occured while marshalising the data into JSON: %w", err)
	}
	// запись
	file.WriteString(string(data))
	return nil
}

func IsValidID(id string) bool {
	for _, r := range id {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func BinarySearchOfTaskID(tasks []Task, id int) int {
	start, end := 0, len(tasks)-1

	for start <= end {
		mid := (start + end) / 2

		if tasks[mid].ID == id {
			return mid
		} else if tasks[mid].ID > id {
			end = mid - 1
		} else {
			start = mid + 1
		}

	}
	return -1
}

func FindTaskID(id string, tasks []Task) (int, error) {
	if !IsValidID(id) {
		return -1, fmt.Errorf("invalid taskID format: %s", id)
	}

	convertedID, err := strconv.Atoi(id)
	if err != nil {
		return -1, fmt.Errorf("failed to convert taskID '%s' to integer: %w", id, err)
	}

	//  Поиск задачи (бинарынй поиск)
	taskIndex := BinarySearchOfTaskID(tasks, convertedID)
	if taskIndex == -1 {
		return -1, fmt.Errorf("task with ID %d not found", convertedID)
	}

	return taskIndex, nil
}

func PrintTask(task Task) {
	fmt.Printf("-TaskID: %v\nDescription: %v\nStatus: %v\ncreatedAt: %v\nupdatedAt: %v\n\n",
		task.ID,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt)
}
