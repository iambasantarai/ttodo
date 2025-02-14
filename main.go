package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func validateID(id int64) {
	if id < 1 {
		log.Fatal("Invalid ID - must be a positive integer")
	}
}

func printTitle() {
	ASCII := "\033[38;5;45;1m  __  __          __   \n" +
		"\033[38;5;45;1m / /_/ /____  ___/ /__ \n" +
		"\033[38;5;45;1m/ __/ __/ _ \\/ _  / _ \\\n" +
		"\033[38;5;45;1m\\__/\\__/\\___/\\_,_/\\___/\033[0m"

	fmt.Printf("%s\n", ASCII)

	fmt.Printf("%s\n", daySay())
}

// Prints usage instructions for the CLI tool.
func helpMenu() {
	fmt.Printf(`
Usage:
  todo <command> [options]

Commands:
  add       -t "Title"           Add a new todo
  toggle    -i ID                Toggle completion status of a todo
  update    -i ID -t "New Title" Update a todo title
  remove    -i ID                Remove a todo
  clean                          Remove completed todos
  list                           Show all todos
  help, --help, -h               Show help menu
`)
}

func handleAdd(store *Store, args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCmd.String("t", "Untitled", "Title for todo")
	addCmd.Parse(args)

	if err := store.AddTodo(*addTitle); err != nil {
		log.Fatal("Failed to add todo:", err)
	}
	fmt.Println("Todo added successfully")
}

func handleToggle(store *Store, args []string) {
	toggleCmd := flag.NewFlagSet("toggle", flag.ExitOnError)
	toggleId := toggleCmd.Int64("i", -1, "ID of todo to toggle status")
	toggleCmd.Parse(args)

	validateID(*toggleId)

	if err := store.ToggleTodo(*toggleId); err != nil {
		log.Fatal("Failed to toggle todo:", err)
	}
	fmt.Println("Todo toggled successfully")
}

func handleUpdate(store *Store, args []string) {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCmd.Int64("i", -1, "ID of todo to update")
	updateTitle := updateCmd.String("t", "Untitled", "New title for todo")
	updateCmd.Parse(args)

	validateID(*updateId)

	if err := store.UpdateTodo(*updateId, *updateTitle); err != nil {
		log.Fatal("Failed to update todo:", err)
	}
	fmt.Println("Todo updated successfully")
}

func handleRemove(store *Store, args []string) {
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	removeId := removeCmd.Int64("i", -1, "ID of todo to remove")
	removeCmd.Parse(args)

	validateID(*removeId)

	if err := store.RemoveTodo(*removeId); err != nil {
		log.Fatal("Failed to remove todo:", err)
	}
	fmt.Println("Todo removed successfully")
}

func handleClean(store *Store) {
	if err := store.Clean(); err != nil {
		log.Fatal("Failed to clean completed todos:", err)
	}
	fmt.Println("Completed todos removed successfully")
}

func handleList(store *Store) {
	todos, err := store.GetTodos()
	if err != nil {
		log.Fatal("Failed to retrieve todos:", err)
	}

	fmt.Println("Todos:")
	for _, todo := range todos {
		status := " "
		if todo.Completed {
			status = "x"
		}
		fmt.Printf("[%s] %-3d: %s\n", status, todo.Id, todo.Title)
	}
}

func main() {
	printTitle()

	store := &Store{}
	if err := store.Init(); err != nil {
		log.Fatalf("unable to init store: %v", err)
	}
	defer store.Close()

	cliArgs := os.Args
	if len(cliArgs) < 2 {
		helpMenu()
		os.Exit(1)
	}

	opCommand := cliArgs[1]
	args := cliArgs[2:]

	switch opCommand {
	case "add":
		handleAdd(store, args)
	case "toggle":
		handleToggle(store, args)
	case "update":
		handleUpdate(store, args)
	case "remove":
		handleRemove(store, args)
	case "clean":
		handleClean(store)
	case "list":
		handleList(store)
	case "help", "--help", "-h":
		helpMenu()
	default:
		fmt.Printf("Unknown command: %s\n", opCommand)
		os.Exit(1)
	}
}
