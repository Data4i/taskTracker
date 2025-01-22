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
		Status:      "todo",
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
	case "in-progress":
		t[idx].Status = "in-progress"
	case "done":
		t[idx].Status = "done"
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

	if len(*tasks) == 1 && (*tasks)[0].Description == "" {
		fmt.Println("No tasks found")
		return
	}

	t := *tasks
	t = append(t[1:2], t[2:]...)

	for idx, task := range t {

		updatedAt := ""

		if task.UpdatedAt != nil {
			updatedAt = task.UpdatedAt.Format(time.RFC1123)
		}

		switch taskStatus {
		case "in-progress":
			if task.Status == "in-progress" {
				taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.CreatedAt.Format(time.RFC1123), updatedAt)
			}
		case "done":
			if task.Status == "done" {
				taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.CreatedAt.Format(time.RFC1123), updatedAt)
			}
		case "todo":
			if task.Status == "todo" {
				taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.CreatedAt.Format(time.RFC1123), updatedAt)
			}
		case "":
			taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.CreatedAt.Format(time.RFC1123), updatedAt)
		default:
			fmt.Println("Invalid task status command")
			return
		}
	}
	taskTable.Render()
}
