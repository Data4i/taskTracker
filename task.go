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

func (tasks *Tasks) UpdateDescription(idx int, description string) error {
	t := *tasks
	if err := t.validateIndex(idx); err != nil {
		return err
	}

	updatedTime := time.Now()
	t[idx].TimeUpdated = &updatedTime
	t[idx].Description = description
	return nil
}

func statusIsPresent(status string) bool {
	statusTypes := []string{
		"todo",
		"in-progress",
		"done",
	}
	present := false
	for _, t := range statusTypes {
		if status == t {
			present = true
		}
	}
	return present

}

func (tasks *Tasks) UpdateStatus(idx int, status string) error {
	t := *tasks
	if err := t.validateIndex(idx); err != nil {
		return err
	}

	isPresent := statusIsPresent(status)
	if isPresent {
		updatedTime := time.Now()
		t[idx].TimeUpdated = &updatedTime

		t[idx].Status = status
	} else {
		return errors.New("invalid status")
	}

	return nil
}

func (tasks *Tasks) Print(status string) {
	taskTable := table.New(os.Stdout)
	taskTable.SetRowLines(false)
	taskTable.SetHeaders("#", "Description", "Status", "TimeCreated", "TimeUpdated")
	for idx, task := range *tasks {
		updatedTime := ""

		if task.TimeUpdated != nil {
			updatedTime = task.TimeUpdated.Format(time.RFC1123)
		}

		switch status {
		case "todo":
			if (*tasks)[idx].Status == "todo" {
				taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.TimeCreated.Format(time.RFC1123), updatedTime)
			}
		case "in-progress":
			if (*tasks)[idx].Status == "in-progress" {
				taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.TimeCreated.Format(time.RFC1123), updatedTime)
			}
		case "done":
			if (*tasks)[idx].Status == "done" {
				taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.TimeCreated.Format(time.RFC1123), updatedTime)
			}
		case "all":
			taskTable.AddRow(strconv.Itoa(idx), task.Description, task.Status, task.TimeCreated.Format(time.RFC1123), updatedTime)
		default:
			fmt.Printf("Status not supported: %s", status)
			return
		}
	}
	taskTable.Render()
}
