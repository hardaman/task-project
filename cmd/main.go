package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"task-project/internal/task"
	"task-project/internal/utils"
)

func main() {
	taskManager := task.NewTaskManager()

	// Load tasks from the file on application start
	tasks, err := utils.LoadTasksFromFile()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if tasks != nil {
		taskManager = task.NewTaskManager()
		taskManager.tasks = tasks
	}

	for {
		printMenu()
		choice := getUserChoice()

		switch choice {
		case 1:
			addTask(taskManager)
		case 2:
			viewTasks(taskManager)
		case 3:
			updateTask(taskManager)
		case 4:
			removeCompletedTasks(taskManager)
		case 5:
			saveTasksToFile(taskManager)
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func printMenu() {
	fmt.Println("GoToDo - A Command-Line Todo List Manager")
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Update Task")
	fmt.Println("4. Remove Completed Tasks")
	fmt.Println("5. Save and Exit")
	fmt.Println("Choose an option (1-5):")
}

func getUserChoice() int {
	var choice int
	fmt.Scanln(&choice)
	return choice
}

func addTask(tm *task.TaskManager) {
	var title, description string
	fmt.Println("Enter task title:")
	fmt.Scanln(&title)
	fmt.Println("Enter task description (optional):")
	fmt.Scanln(&description)

	tm.AddTask(title, description)
	fmt.Println("Task added successfully!")
}

func viewTasks(tm *task.TaskManager) {
	fmt.Println("Current Tasks:")
	tm.ViewTasks()
}

func updateTask(tm *task.TaskManager) {
	viewTasks(tm)
	fmt.Println("Enter the task number you want to update:")
	index := getUserChoice() - 1

	if index < 0 || index >= len(tm.tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	var title, description string
	var completed bool

	fmt.Println("Enter new task title:")
	fmt.Scanln(&title)
	fmt.Println("Enter new task description (optional):")
	fmt.Scanln(&description)
	fmt.Println("Is the task completed? (true/false):")
	input := strings.ToLower(getUserInput())

	if input == "true" {
		completed = true
	}

	tm.UpdateTask(index, title, description, completed)
	fmt.Println("Task updated successfully!")
}

func removeCompletedTasks(tm *task.TaskManager) {
	tm.RemoveCompletedTasks()
	fmt.Println("Completed tasks removed successfully!")
}

func saveTasksToFile(tm *task.TaskManager) {
	err := utils.SaveTasksToFile(tm.tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Println("Tasks saved to file.")
}

func getUserInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}
