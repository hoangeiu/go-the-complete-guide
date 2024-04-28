package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type Saver interface {
	Save() error
}

type OutputTable interface {
	Display()
	Saver
}

func main() {
	// Type any
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")
	// Generics
	result := add(1, 2)
	fmt.Println(result)

	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
}

func outputData(data OutputTable) error {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

// Type any
func printSomething(value any) {
	// Check type (1 type only)
	typedVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", typedVal)
		return
	}

	// Check type (multi type)
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// }
}

// Generics
func add[T int | float64](a, b T) T {
	return a + b
}
