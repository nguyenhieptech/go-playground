package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type ToDoList struct {
	Tasks []Task `json:"tasks"`
}

const dataFile = "tasks.json"

func main() {
	fmt.Println("Welcome to CLI Todo App")

	todoList := loadTasks()

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Incomplete Task")
		fmt.Println("6. Exit")

		fmt.Print("Enter your choice: ")
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addTask(&todoList)
		case "2":
			listTasks(&todoList)
		case "3":
			completeTask(&todoList)
		case "4":
			deleteTask(&todoList)
		case "5":
			incompleteTask(&todoList)
		case "6":
			fmt.Println("Exiting... Goodbye!")
			saveTasks(&todoList)
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func addTask(todoList *ToDoList) {
	fmt.Print("Enter task description: ")
	reader := bufio.NewReader(os.Stdin)
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	newTask := Task{
		ID:          len(todoList.Tasks) + 1,
		Description: description,
		Completed:   false,
	}

	todoList.Tasks = append(todoList.Tasks, newTask)
	saveTasks(todoList)

	fmt.Println("Task added!")
}

func listTasks(todoList *ToDoList) {
	if len(todoList.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("\nTasks:")
	for _, task := range todoList.Tasks {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", task.ID, status, task.Description)
	}
}

func completeTask(todoList *ToDoList) {
	fmt.Print("Enter task ID to mark as completed: ")
	reader := bufio.NewReader(os.Stdin)
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	for i := range todoList.Tasks {
		if todoList.Tasks[i].ID == id {
			todoList.Tasks[i].Completed = true
			saveTasks(todoList)
			fmt.Println("Task marked as completed!")
			return
		}
	}

	fmt.Println("Task not found.")
}

func deleteTask(todoList *ToDoList) {
	fmt.Print("Enter task ID to delete: ")
	reader := bufio.NewReader(os.Stdin)
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	for i, task := range todoList.Tasks {
		if task.ID == id {
			todoList.Tasks = append(todoList.Tasks[:i], todoList.Tasks[i+1:]...)
			saveTasks(todoList)
			fmt.Println("Task deleted!")
			return
		}
	}

	fmt.Println("Task not found.")
}

func incompleteTask(todoList *ToDoList) {
	fmt.Print("Enter task ID to mark as uncompleted: ")
	reader := bufio.NewReader(os.Stdin)
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	for i := range todoList.Tasks {
		if todoList.Tasks[i].ID == id {
			todoList.Tasks[i].Completed = false
			fmt.Println("Task marked as uncompleted!")
			saveTasks(todoList)
			return
		}
	}

	fmt.Println("Task not found.")
}

func loadTasks() ToDoList {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return ToDoList{}
		}
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}
	defer file.Close()

	var todoList ToDoList
	err = json.NewDecoder(file).Decode(&todoList)
	if err != nil {
		fmt.Println("Error decoding tasks:", err)
	}

	return todoList
}

func saveTasks(todoList *ToDoList) {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(todoList)
	if err != nil {
		fmt.Println("Error encoding tasks:", err)
	}
}
