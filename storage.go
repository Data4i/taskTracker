package main

import (
	"encoding/json"
	"os"
	"time"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		taskData := Task{
			Description: "",
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		}
		initialTasks := []Task{taskData}

		fileData, jsonErr := json.MarshalIndent(initialTasks, "", "    ")

		if jsonErr != nil {
			panic(jsonErr)
		}

		if createErr := os.WriteFile(fileName, fileData, 0644); createErr != nil {
			panic(createErr)
		}

	}
	return &Storage[T]{fileName}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, data)
}
