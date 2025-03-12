package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID   int
	Name string
	Done bool
}

func main() {
	tasks := []Task{}
	index := 1
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a task (or type 'Exit' to stop):")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "Exit" {
			break
		}

		tasks = append(tasks, Task{ID: index, Name: input})
		index++
	}

	
	fmt.Println("\nTodo List:")
	for _, task := range tasks {
		fmt.Printf("%d. Name: %s\n", task.ID, task.Name)
	}
}