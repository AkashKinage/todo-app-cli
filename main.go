package main

import (
	"fmt"
	"os"
)

type Task struct {
	ID int
	Name string
	Completed bool
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Usage:\n\ntodo add\ntodo list\ntodo delete")
		return
	}

	choice := os.Args[1]

	switch choice {
	case "add":
		fmt.Println("Adding task...")
	case "list":
		fmt.Println("Listing tasks...")
	case "delete":
		fmt.Println("Deleting task...")
	default:
		fmt.Println("Invalid choice")
	}

	task := Task{
		ID: 1,
		Name: "Task 1",
		Completed: false,
	}

	fmt.Printf("Task: %+v\n", task)
}