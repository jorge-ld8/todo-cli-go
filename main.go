package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jorge-ld8/todo-cli/todo"
)

func main() {
	// three main parts

	// if len < 2 exit (general error checking)
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// load todos or try to, if error exit
	todos := &todo.Todos{}
	if err := todos.Load(); err != nil {
		fmt.Println("Error loading todos:", err)
		os.Exit(1)
	}
	// switch on os.args[1]
	switch os.Args[1] {
	case "add":
		cmdAdd(todos, os.Args[2:])
	case "list":
		cmdList(todos)
	default:
		printUsage()
		os.Exit(1)
	}
}

// cmd add
func cmdAdd(todos *todo.Todos, args []string) {
	if len(args) == 0 {
		fmt.Println("Provide a title for the todo item.")
		os.Exit(1)
	}

	title := strings.Join(args, " ")
	item := todos.Add(title)

	if err := todos.Save(); err != nil {
		fmt.Println("Error saving todo:", err)
		os.Exit(1)
	}

	fmt.Printf("Added todo: %d - %s\n", item.ID, item.Title)
}

// cmd list
func cmdList(todos *todo.Todos) {
	// bring the list of todos (from todo/todo.go)
	items := todos.List()

	if len(items) == 0 {
		fmt.Println("No todo items added. Add one with 'todo add <task>'")
		return
	}

	fmt.Println("\nTodo List:")
	fmt.Println(strings.Repeat("-", 40))

	// iterate over todos list
	for _, item := range items {
		status := "[ ]"
		if item.Completed {
			status = "[✓]"
		}
		fmt.Printf("%s %d: %s\n", status, item.ID, item.Title)
	}
	fmt.Println()
}

// if you typed an unknown command print usage manual
func printUsage() {
	fmt.Println(`Todo CLI - A simple task manager
Usage:
  todo <command> [arguments]

Commands:
  add <title>      Add a new todo
  list             List all todos`)
}
