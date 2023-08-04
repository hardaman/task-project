package task

import (
	"fmt"
)

type TaskManager struct {
	tasks []*Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make([]*Task, 0),
	}
}

func (tm *TaskManager) AddTask(title, description string) {
	task := &Task{
		Title:       title,
		Description: description,
		Completed:   false,
	}
	tm.tasks = append(tm.tasks, task)
}

func (tm *TaskManager) ViewTasks() {
	for i, task := range tm.tasks {
		status := "Not Completed"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s (%s)\n", i+1, task.Title, status)
	}
}

func (tm *TaskManager) UpdateTask(index int, title, description string, completed bool) {
	if index >= 0 && index < len(tm.tasks) {
		task := tm.tasks[index]
		task.Title = title
		task.Description = description
		task.Completed = completed
	}
}

func (tm *TaskManager) RemoveCompletedTasks() {
	var updatedTasks []*Task
	for _, task := range tm.tasks {
		if !task.Completed {
			updatedTasks = append(updatedTasks, task)
		}
	}
	tm.tasks = updatedTasks
}
