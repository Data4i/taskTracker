package main

import (
	"errors"
	"fmt"
	"github.com/aquasecurity/table"
	"os"
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

func (tasks *Tasks) updateStatusInProgress(idx int) error {
	t := *tasks
	if err := t.validateTask(idx); err != nil {
		return err
	}
	if t[idx].Status == "In Progress" {
		fmt.Println("Task is already In Progress")
	}

	t[idx].Status = "In Progress"
	fmt.Printf("Task %s is now In Progress", t[idx].Description)
	return nil
}

func (tasks *Tasks) updateStatusDone(idx int) error {
	t := *tasks
	if err := t.validateTask(idx); err != nil {
		return err
	}
	if t[idx].Status == "Done" {
		fmt.Println("Task is already done")
	}

	t[idx].Status = "Done"
	fmt.Printf("Task %s is done", t[idx].Description)
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
		taskTable.AddRow(idx, task.Description, task.Status, task.CreatedAt.Format(time.RFC1123), updatedAt)
	}
}
