package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Title     string
	Completed bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := make([]Task, 0)

	fmt.Println("Terminal To-Do App")
	fmt.Println("-------------------")

	for {
		printMenu()
		choice := readLine(scanner, "Choose an option: ")

		switch choice {
		case "1":
			addTask(scanner, &tasks)
		case "2":
			listTasks(tasks)
		case "3":
			markTaskDone(scanner, &tasks)
		case "4":
			deleteTask(scanner, &tasks)
		case "5":
			fmt.Println("Goodbye.")
			return
		default:
			fmt.Println("Invalid option. Enter 1-5.")
		}
	}
}

func printMenu() {
	fmt.Println()
	fmt.Println("1) Add task")
	fmt.Println("2) List tasks")
	fmt.Println("3) Mark task as done")
	fmt.Println("4) Delete task")
	fmt.Println("5) Exit")
}

func readLine(scanner *bufio.Scanner, prompt string) string {
	for {
		fmt.Print(prompt)
		if scanner.Scan() {
			return strings.TrimSpace(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Input error:", err)
		}
		os.Exit(1)
	}
}

func addTask(scanner *bufio.Scanner, tasks *[]Task) {
	title := readLine(scanner, "Task title: ")
	if title == "" {
		fmt.Println("Task title cannot be empty.")
		return
	}

	*tasks = append(*tasks, Task{Title: title})
	fmt.Println("Task added.")
}

func listTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet.")
		return
	}

	fmt.Println("Tasks:")
	for i, task := range tasks {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, task.Title)
	}
}

func markTaskDone(scanner *bufio.Scanner, tasks *[]Task) {
	if len(*tasks) == 0 {
		fmt.Println("No tasks to mark.")
		return
	}

	listTasks(*tasks)
	idx, ok := readTaskNumber(scanner, len(*tasks), "Enter task number to mark done: ")
	if !ok {
		return
	}

	if (*tasks)[idx].Completed {
		fmt.Println("Task is already completed.")
		return
	}

	(*tasks)[idx].Completed = true
	fmt.Println("Task marked as done.")
}

func deleteTask(scanner *bufio.Scanner, tasks *[]Task) {
	if len(*tasks) == 0 {
		fmt.Println("No tasks to delete.")
		return
	}

	listTasks(*tasks)
	idx, ok := readTaskNumber(scanner, len(*tasks), "Enter task number to delete: ")
	if !ok {
		return
	}

	*tasks = append((*tasks)[:idx], (*tasks)[idx+1:]...)
	fmt.Println("Task deleted.")
}

func readTaskNumber(scanner *bufio.Scanner, max int, prompt string) (int, bool) {
	input := readLine(scanner, prompt)
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Please enter a valid number.")
		return 0, false
	}

	if n < 1 || n > max {
		fmt.Printf("Please enter a number between 1 and %d.\n", max)
		return 0, false
	}

	return n - 1, true
}
