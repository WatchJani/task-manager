package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     time.Time
}

var (
	tasks   []*Task
	taskMap map[int]*Task
)

func init() {
	taskMap = make(map[int]*Task)
}

func createTask(title, description string, dueDate time.Time) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("Title is empty\n")
	}

	newTask := &Task{
		ID:          len(tasks) + 1,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
	}

	taskMap[len(tasks)] = newTask
	tasks = append(tasks, newTask)

	log.Printf("Task is created\n")
	return newTask, nil
}

func updateTask(taskID int, title, description string, dueDate time.Time) error {
	task, err := findTaskByID(taskID)

	if err != nil {
		return fmt.Errorf(fmt.Sprintf("ID %d is not exist\n", taskID))
	}

	if title == "" {
		return fmt.Errorf("Title is empty\n")
	}

	*task = Task{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
	}

	log.Println(fmt.Sprintf("Task %d is updated", taskID))
	return nil
}

func deleteTask(taskID int) error {
	_, err := findTaskByID(taskID)

	if err != nil {
		return fmt.Errorf(fmt.Sprintf("ID %d is not exist\n", taskID))
	}

	tasks = append(tasks[:taskID], tasks[taskID+1:]...)
	log.Println(fmt.Sprintf("Task %d is deleted", taskID))

	return nil
}

func getTask(taskID int) (Task, error) {
	return *tasks[taskID], nil
}

func getAllTasks() ([]*Task, error) {
	return tasks, nil
}

func findTaskByID(taskID int) (*Task, error) {
	task, found := taskMap[taskID]
	if !found {
		return nil, errors.New("Task not found")
	}
	return task, nil
}

func searchTasks(keyword string) ([]Task, error) {
	searched := []Task{}

	for _, task := range tasks {
		if containsKeyword(task, keyword) {
			searched = append(searched, *task)
		}
	}

	return searched, nil
}

func containsKeyword(task *Task, keyword string) bool {
	if keyword == "" {
		return true
	}

	return containsString(task.Title, keyword) || containsString(task.Description, keyword)
}

func containsString(str, substr string) bool {
	return len(str) >= len(substr) && str[:len(substr)] == substr
}

func main() {
	createTask("Janko", "Top", time.Now().AddDate(0, 0, 7))
	updateTask(0, "super", "asd", time.Now())
	deleteTask(0)

	allTasks, _ := getAllTasks()
	fmt.Println(allTasks)
}
