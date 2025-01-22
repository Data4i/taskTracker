package main

import (
	"fmt"
	"github.com/aquasecurity/table"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Description string
	Status      string
	TimeCreated time.Time
	TimeUpdated *time.Time
}

type Tasks []Task

func (tasks *Tasks) Add(description string) {
	task := Task{
		Description: description,
		Status:      "todo",
		TimeCreated: time.Now(),
		TimeUpdated: nil,
	}

	*tasks = append(*tasks, task)
}

func (tasks *Tasks) validateIndex(idx int) error {
	if idx < 0 || idx >= len(*tasks) {
		err := fmt.Errorf("index out of range")
		fmt.Println(err)
		return err
	}
	return nil
}

func (tasks *Tasks) Delete(idx int) error {
	t := *tasks
	if err := t.validateIndex(idx); err != nil {
		return err
	}
	*tasks = append(t[:idx], t[idx+1:]...)
	return nil
}

func (tasks *Tasks) Print() {
	taskTable := table.New(os.Stdout)
	taskTable.SetRowLines(false)
	taskTable.SetHeaders("#", "Description", "Status", "TimeCreated", "TimeUpdated")

	for idx, task := range *tasks {
		updatedTime := ""

		if task.TimeUpdated != nil {
			updatedTime = task.TimeUpdated.Format(time.RFC1123)
		}

		taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.TimeCreated.Format(time.RFC1123), updatedTime)
	}
	taskTable.Render()
}
