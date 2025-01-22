package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{FileName: filename}
}

func (storage *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(storage.FileName, fileData, 0644)
}

func (storage *Storage[T]) Load(data *T) error {
	if _, err := os.Stat(storage.FileName); os.IsNotExist(err) {
		return nil
	}
	fileData, err := os.ReadFile(storage.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
