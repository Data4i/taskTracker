package main

import (
	"errors"
	"fmt"
	"github.com/aquasecurity/table"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type Tasks []Task

func (tasks *Tasks) add(description string) {
	task := Task{
		Description: description,
		Status:      "In Progress",
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}
	*tasks = append(*tasks, task)
}

func (tasks *Tasks) validateTask(idx int) error {
	if idx < 0 || idx >= len(*tasks) {
		return errors.New("invalid task index")
	}
	return nil
}

func (tasks *Tasks) updateDescription(idx int, description string) error {
	t := *tasks
	if err := t.validateTask(idx); err != nil {
		return err
	}
	t[idx].Description = description
	return nil
}

func (tasks *Tasks) updateStatus(idx int, status string) error {
	t := *tasks

	if err := t.validateTask(idx); err != nil {
		return err
	}

	if t[idx].Status == status {
		fmt.Printf("Task is aready %s", status)
		return nil
	}

	switch status {
	case "In Progress":
		t[idx].Status = "Done"
	case "Done":
		t[idx].Status = "In Progress"
	}

	return nil
}

func (tasks *Tasks) delete(idx int) error {
	t := *tasks
	if err := t.validateTask(idx); err != nil {
		return err
	}
	*tasks = append(t[:idx], t[idx+1:]...)
	return nil
}

func (tasks *Tasks) print() {
	taskTable := table.New(os.Stdout)
	taskTable.SetRowLines(false)
	taskTable.SetHeaders("#", "Description", "Status", "CreatedAt", "UpdatedAt")

	for idx, task := range *tasks {

		updatedAt := ""
		if task.UpdatedAt != nil {
			updatedAt = task.UpdatedAt.Format(time.RFC1123)
		}
		taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.CreatedAt.Format(time.RFC1123), updatedAt)
	}
}
