package main

import (
	"fmt"
	"os"
)

var globalTaskId = 0

type Task struct {
	ID int
	Name string
	Completed bool
}

var tasks []Task

func addTask(name string) {
	task := Task{
		ID: globalTaskId + 1,
		Name: name,
		Completed: false,
	}
	tasks = append(tasks, task)
	fmt.Printf("Task added: %+v\n", task)
	globalTaskId++
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Usage:\n\ntodo add\ntodo list\ntodo delete")
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