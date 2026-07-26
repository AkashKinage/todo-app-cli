package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
}

var tasks []Task

func getLastTaskId() int {
	if len(tasks) == 0 {
		return 0
	}
	return tasks[len(tasks)-1].ID
}

func addTask(name string) {
	task := Task{
		ID: getLastTaskId() + 1,
		Name: name,
		Completed: false,
	}
	tasks = append(tasks, task)
	tasksJSON, _ := json.Marshal(tasks)
	err := os.WriteFile("tasks.json", tasksJSON, 0644)
	if err != nil {
		fmt.Println("Error writing tasks file:", err)
		return
	}
	fmt.Printf("Task added: %+v\n", task)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Usage:\n\ntodo add\ntodo list\ntodo delete")
		return
	}

	data, err := os.ReadFile("tasks.json")
	if err == nil {
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println("Error unmarshaling tasks:", err)
			return
		}
	} else if os.IsNotExist(err) {
		fmt.Println("Tasks file not found. Starting with an empty task list.")
		os.WriteFile("tasks.json", []byte("[]"), 0644)
	} else {
		fmt.Println("Error reading tasks file:", err)
		return
	}

	choice := os.Args[1]

	switch choice {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task name")
			return
		}

		addTask(os.Args[2])
	case "list":
		fmt.Println("Listing tasks...")
	case "delete":
		fmt.Println("Deleting task...")
	default:
		fmt.Println("Invalid choice")
	}
}