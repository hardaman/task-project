package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"task-project/internal/task"
)

const dataFileName = "tasks.json"

func SaveTasksToFile(tasks []*task.Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	filePath := getFilePath()
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadTasksFromFile() ([]*task.Task, error) {
	filePath := getFilePath()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, nil
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var tasks []*task.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func getFilePath() string {
	currentDir, _ := os.Getwd()
	return filepath.Join(currentDir, dataFileName)
}
