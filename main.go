package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const taskFile = "tasks.json"

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var tasks []Task

func getLastTaskID() int {
	if len(tasks) == 0 {
		return 0
	}
	return tasks[len(tasks)-1].ID
}

func writeTasksToFile() error {
	tasksJSON, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(taskFile, tasksJSON, 0644)
	if err != nil {
		return err
	}
	return nil
}

func parseTaskID(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func findTaskByID(id int) (*Task, error) {
	// NOTE: The iteration variable `task` in a `range` loop is a copy of the slice element, not the actual element stored in `tasks`.
	// Any changes made to `tasks[i]` are not reflected in `task`, so printing `task` after updating `tasks[i]` will show the old value.
	// If you need to modify or return the original task, work with `tasks[i]` (or `&tasks[i]` to get a pointer) instead of the `task` variable.

	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i], nil
		}
	}

	return nil, errors.New("no task found")
}

func findTaskIndexByID(id int) (int, error) {
	for ind, task := range tasks {
		if task.ID == id {
			return ind, nil
		}
	}

	return 0, errors.New("no task found")
}

func addTask(name string) {
	task := Task{
		ID:        getLastTaskID() + 1,
		Name:      name,
		Completed: false,
	}
	tasks = append(tasks, task)
	if err := writeTasksToFile(); err != nil {
		fmt.Println("Error writing tasks to file:", err)
		return
	}
	fmt.Printf("Task added: %+v\n", task)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("%d. %s %s\n", task.ID, status, task.Name)
	}
}

func completeTask(idStr string) {
	id, err := parseTaskID(idStr)
	if err != nil {
		fmt.Println("Invalid task ID", err)
		return
	}

	task, err := findTaskByID(id)
	if err != nil {
		fmt.Println("Task not found:", err)
		return
	}

	task.Completed = true
	if err := writeTasksToFile(); err != nil {
		fmt.Println("Error writing tasks to file:", err)
		return
	}

	fmt.Println("Task completed:", *task)
}

func deleteTask(idStr string) {
	id, err := parseTaskID(idStr)
	if err != nil {
		fmt.Println("Invalid task ID", err)
		return
	}

	idx, err := findTaskIndexByID(id)
	if err != nil {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}

	deletedTask := tasks[idx]

	tasks = append(tasks[:idx], tasks[idx+1:]...)
	if err := writeTasksToFile(); err != nil {
		fmt.Println("Error writing tasks to file:", err)
		return
	}

	fmt.Printf("Task deleted: %+v\n", deletedTask)
}

func loadTasks() error {
	data, err := os.ReadFile(taskFile)
	if err == nil {
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println("Error unmarshaling tasks:", err)
			return err
		}
	} else if os.IsNotExist(err) {
		fmt.Println("Tasks file not found. Starting with an empty task list.")
		if err := os.WriteFile(taskFile, []byte("[]"), 0644); err != nil {
			fmt.Println("Error creating tasks file:", err)
			return err
		}
	} else {
		fmt.Println("Error reading tasks file:", err)
		return err
	}

	return nil
}

func handleChoice(choice string) {
	switch choice {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task name")
			return
		}

		name := strings.Join(os.Args[2:], " ")
		addTask(name)
	case "list":
		listTasks()
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID to complete")
			return
		}

		completeTask(os.Args[2])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID to delete")
			return
		}

		deleteTask(os.Args[2])
	default:
		fmt.Println("Invalid choice")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Usage:\n\ntodo add\ntodo list\ntodo complete\ntodo delete")
		return
	}

	if err := loadTasks(); err != nil {
		fmt.Println("Error loading Tasks file:", err)
		return
	}

	handleChoice(os.Args[1])
}
