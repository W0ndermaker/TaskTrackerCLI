package internal

import (
	"encoding/json"
	"os"
	"testing"
)

func TestReadDataFromStorage(t *testing.T) {
	tempFile, err := os.CreateTemp("", "storage_*.go")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tempFile.Name())

	tmpFileName := tempFile.Name()

	t.Run("File doesn't exists", func(t *testing.T) {
		os.Remove(tmpFileName)

		tasks, err := ReadDataFromStorage(tmpFileName)
		if err != nil {
			t.Errorf("Expected no errors, got: %v", err)
		}
		if tasks == nil {
			t.Errorf("Expected empty slice, got nil")
		}
		if len(tasks) != 0 {
			t.Errorf("Expected empty slice, got %v items", len(tasks))
		}
	})

	t.Run("file is empty", func(t *testing.T) {
		err := os.WriteFile(tmpFileName, []byte{}, 0644)
		if err != nil {
			t.Fatal(err)
		}

		tasks, err := ReadDataFromStorage(tmpFileName)
		if err != nil {
			t.Errorf("Expected no errors, got: %v", err)
		}
		if len(tasks) != 0 {
			t.Errorf("Expected empty slice, got %v items", len(tasks))
		}
	})

	t.Run("valid json data", func(t *testing.T) {
		testTasks := []Task{
			{ID: 1, Description: "task1"},
			{ID: 2, Description: "task2"},
			{ID: 3, Description: "task3"},
		}

		data, err := json.Marshal(testTasks)
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile(tmpFileName, data, 0644)
		if err != nil {
			t.Fatal(err)
		}

		tasks, err := ReadDataFromStorage(tmpFileName)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(tasks) != len(testTasks) {
			t.Errorf("Expected %d tasks, got %d", len(testTasks), len(tasks))
		}

		for i, task := range tasks {
			if task.ID != testTasks[i].ID || task.Description != testTasks[i].Description {
				t.Errorf("Task mismatch at index %d", i)
			}
		}

	})

	t.Run("invalid json data", func(t *testing.T) {
		err = os.WriteFile(tmpFileName, []byte(`invalid data`), 0644)
		if err != nil {
			t.Fatal(err)
		}

		_, err := ReadDataFromStorage(tmpFileName)
		if err == nil {
			t.Error("Expected error for invalid json, got nil")
		}
	})

}

func TestWriteToStorage(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test_storage_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())
	tmpFileName := tempFile.Name()
	t.Run("Write to new file", func(t *testing.T) {

		testTasks := []Task{
			{ID: 1, Description: "Task 1"},
			{ID: 2, Description: "Task 2"},
		}

		err := WriteToStorage(tmpFileName, testTasks)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		data, err := os.ReadFile(tmpFileName)
		if err != nil {
			t.Fatal(err)
		}

		var tasks []Task
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			t.Fatal(err)
		}

		if len(tasks) != len(testTasks) {
			t.Errorf("Expected %d tasks, got %d", len(testTasks), len(tasks))
		}
		os.Remove(tmpFileName)
	})

	t.Run("empty tasks slice", func(t *testing.T) {
		tasks := []Task{}

		err := WriteToStorage(tmpFileName, tasks)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		data, err := os.ReadFile(tmpFileName)
		if err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}

		expected := "[]"
		if string(data) != expected && string(data) != "[]\n" {
			t.Errorf("Expected empty array, got %s", string(data))
		}
		os.Remove(tmpFileName)
	})
}

func TestIsValid(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  bool
	}{
		{"1 should be true(valid)", "1", true},
		{"5 should be true(valid)", "15", true},
		{"-9 should be false(invalid)", "-9", false},
		{"1144 should be true(valid)", "1144", true},
		{"1aa should be false(invalid)", "1aa", false},
		{"a99 should be false(invalid)", "a99", false},
		{"4.1 should be false(invalid)", "4.1", false},
		{"14,1 should be false(invalid)", "14,1", false},
		{"1 9 should be false(invalid)", "1 9", false},
		{"1 2   3   4 should be false(invalid)", "1 2   3   4", false},
		{"'empty'  should be false(invalid)", " ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := IsValidID(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

}

func TestBinarySearchOfTaskID(t *testing.T) {
	var tests = []struct {
		name  string
		tasks []Task
		id    int
		want  int
	}{
		{"correct performance",
			[]Task{
				{ID: 1, Description: "task1"},
				{ID: 2, Description: "task2"},
				{ID: 3, Description: "task3"},
				{ID: 4, Description: "task4"},
				{ID: 5, Description: "task5"},
			},
			4, 3},
		{"incorrect performance",
			[]Task{
				{ID: 1, Description: "task1"},
				{ID: 2, Description: "task2"},
				{ID: 3, Description: "task3"},
				{ID: 4, Description: "task4"},
				{ID: 5, Description: "task5"},
			},
			6, -1},
		{"empty tasks slice", []Task{}, 1, -1},
		{"one item in tasks slice", []Task{{ID: 2, Description: "task2"}}, 2, 0},
		{"two item in tasks slice",
			[]Task{
				{ID: 2, Description: "task2"},
				{ID: 8, Description: "task8"},
			},
			8, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := BinarySearchOfTaskID(tt.id, tt.tasks)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
