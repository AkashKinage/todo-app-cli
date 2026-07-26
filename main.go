package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func writeTasksToFile() {
	tasksJSON, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error marshaling tasks:", err)
		return
	}
	err = os.WriteFile("tasks.json", tasksJSON, 0644)
	if err != nil {
		fmt.Println("Error writing tasks file:", err)
		return
	}
}

func addTask(name string) {
	task := Task{
		ID: getLastTaskId() + 1,
		Name: name,
		Completed: false,
	}
	tasks = append(tasks, task)
	writeTasksToFile()
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
	var id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID", err)
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			writeTasksToFile()
			// NOTE: `task` from `range` is a copy, not a reference to tasks[i].
			// Mutating tasks[i] doesn't update `task` — print/use tasks[i] instead
			// if you need the updated value after modification.
			// earlier, I mistakenly printed `task` instead of `tasks[i]`, which is a copy and doesn't reflect the updated state.
			fmt.Printf("Task completed: %+v\n", tasks[i])
			return
		}
	}
	fmt.Println("ID is invalid.")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Usage:\n\ntodo add\ntodo list\ntodo complete\ntodo delete")
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
		listTasks()
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID to complete")
			return
		}

		completeTask(os.Args[2])
	case "delete":
		fmt.Println("Deleting task...")
	default:
		fmt.Println("Invalid choice")
	}
}