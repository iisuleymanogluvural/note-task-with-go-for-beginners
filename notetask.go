package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func task1() {
	fmt.Println("Task 1: Create a note")
	fmt.Print("Enter your note: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	note := scanner.Text()
	fmt.Printf("Note created: %s\n", note)
}
func task2() {
	fmt.Println("Task 2: List all notes")
	// In a real application, you would retrieve notes from a database or file.
	// Here we just simulate with a static list.
	notes := []string{"Note 1: Buy groceries", "Note 2: Call mom", "Note 3: Finish project"}
	for _, note := range notes {
		fmt.Println(note)
	}
}
func task3() {
	fmt.Println("Task 3: Delete a note")
	fmt.Print("Enter the note number to delete (1-3): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	noteNumber := scanner.Text()
	// In a real application, you would delete the note from a database or file.
	fmt.Printf("Note %s deleted.\n", noteNumber)
}
func main() {
	fmt.Println("Welcome to the Note Task Application!")
	fmt.Println("Please choose a task:")
	fmt.Println("1. Create a note")
	fmt.Println("2. List all notes")
	fmt.Println("3. Delete a note")
	fmt.Print("Enter your choice (1-3): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()

	switch choice {
	case "1":
		task1()
	case "2":
		task2()
	case "3":
		task3()
	default:
		fmt.Println("Invalid choice. Please try again.")
	}

	time.Sleep(2 * time.Second) // Pause for 2 seconds before exiting
}
