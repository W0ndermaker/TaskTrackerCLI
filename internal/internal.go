package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// reads data from storage file if it exists else create a new one
func ReadDataFromStorage(filename string) ([]Task, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error occured while opening storage file:")
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()

	if fileSize == 0 {
		return []Task{}, nil
	}

	myBuf := make([]byte, fileSize)

	_, err = file.Read(myBuf)
	if err != nil {
		fmt.Println("Error while storing read data from file")
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

// writes new tasks to storage file
func WriteToStorage(filename string, newInfo []Task) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Error occured while opening storage file: %w", err)
	}
	defer file.Close()
	data, err := json.MarshalIndent(newInfo, "", " ")
	if err != nil {
		return fmt.Errorf("Error occured while marshalising the data into JSON: %w", err)
	}
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

// Binary search
func BinarySearchOfTaskID(id int, tasks []Task) int {
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

// searches a task with input id and returns index of the task
func FindTaskID(id string, tasks []Task) (int, error) {
	if !IsValidID(id) {
		return -1, fmt.Errorf("invalid taskID format: %s", id)
	}

	convertedID, err := strconv.Atoi(id)
	if err != nil {
		return -1, fmt.Errorf("failed to convert taskID '%s' to integer: %w", id, err)
	}
	taskIndex := BinarySearchOfTaskID(convertedID, tasks)
	if taskIndex == -1 {
		return -1, fmt.Errorf("task with ID %d not found", convertedID)
	}

	return taskIndex, nil
}

// Prints a task
func PrintTask(task Task) {
	fmt.Printf("-TaskID: %v\nDescription: %v\nStatus: %v\ncreatedAt: %v\nupdatedAt: %v\n\n",
		task.ID,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt)
}
