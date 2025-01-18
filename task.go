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

func (tasks *Tasks) Add(description string) {
	task := Task{
		Description: description,
		Status:      "Todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}
	*tasks = append(*tasks, task)
}

func (tasks *Tasks) ValidateTask(idx int) error {
	if idx < 0 || idx >= len(*tasks) {
		return errors.New("invalid task index")
	}
	return nil
}

func (tasks *Tasks) UpdateDescription(idx int, description string) error {
	t := *tasks
	if err := t.ValidateTask(idx); err != nil {
		return err
	}

	updateTime := time.Now()
	t[idx].UpdatedAt = &updateTime

	t[idx].Description = description
	return nil
}

func (tasks *Tasks) UpdateStatus(idx int, status string) error {
	t := *tasks

	updateTime := time.Now()
	t[idx].UpdatedAt = &updateTime

	if err := t.ValidateTask(idx); err != nil {
		return err
	}

	if t[idx].Status == status {
		fmt.Printf("Task is aready %s", status)
		return nil
	}

	switch status {
	case "In Progress":
		t[idx].Status = "In Progress"
	case "Done":
		t[idx].Status = "Done"
	}

	return nil
}

func (tasks *Tasks) delete(idx int) error {
	t := *tasks
	if err := t.ValidateTask(idx); err != nil {
		return err
	}
	*tasks = append(t[:idx], t[idx+1:]...)
	return nil
}

func (tasks *Tasks) print(taskStatus string) {
	taskTable := table.New(os.Stdout)
	taskTable.SetRowLines(false)
	taskTable.SetHeaders("#", "Description", "Status", "CreatedAt", "UpdatedAt")

	for idx, task := range *tasks {
		updatedAt := ""

		if task.UpdatedAt != nil {
			updatedAt = task.UpdatedAt.Format(time.RFC1123)
		}

		switch taskStatus {
		case "In Progress":
			if task.Status == "In Progress" {
				taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.CreatedAt.Format(time.RFC1123), updatedAt)
			}
		case "Done":
			if task.Status == "Done" {
				taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.CreatedAt.Format(time.RFC1123), updatedAt)
			}
		case "Todo":
			if task.Status == "Todo" {
				taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.CreatedAt.Format(time.RFC1123), updatedAt)
			}
		case "":
			taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.CreatedAt.Format(time.RFC1123), updatedAt)
		default:
			fmt.Println("Invalid task status command")
		}
	}
	taskTable.Render()
}
