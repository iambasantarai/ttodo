package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	store := &Store{}
	if err := store.Init(); err != nil {
		log.Fatalf("unable to init store: %v", err)
	}
	defer store.Close()

	cliArgs := os.Args
	if len(cliArgs) < 2 {
		fmt.Println(
			"expected 'add', 'update', 'remove', 'toggle', 'list', or 'clean' subcommand",
		)
		os.Exit(1)
	}

	opCommand := cliArgs[1]

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCmd.String("t", "Untitled", "Title for todo")

	toggleCmd := flag.NewFlagSet("toggle", flag.ExitOnError)
	toggleId := toggleCmd.Int64("i", -1, "ID of todo to toggle status")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCmd.Int64("i", -1, "ID of todo to update")
	updateTitle := updateCmd.String("t", "Untitled", "New title for todo")

	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	removeId := removeCmd.Int64("i", -1, "ID of todo to remove")

	cleanCmd := flag.NewFlagSet("clean", flag.ExitOnError)

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	cliArgs = cliArgs[2:]

	switch opCommand {
	case "add":
		addCmd.Parse(os.Args[2:])
		if err := store.AddTodo(*addTitle); err != nil {
			log.Fatal("Failed to add todo: ", err)
		}
		fmt.Println("Todo added successfully")

	case "toggle":
		toggleCmd.Parse(os.Args[2:])
		if *toggleId < 1 {
			log.Fatal("Invalid ID - must be positive integer")
		}
		if err := store.ToggleTodo(*toggleId); err != nil {
			log.Fatal("Failed to toggle todo:", err)
		}
		fmt.Println("Todo toggled successfully")

	case "update":
		updateCmd.Parse(os.Args[2:])
		if *updateId < 1 {
			log.Fatal("Invalid ID - must be positive integer")
		}
		if err := store.UpdateTodo(*updateId, *updateTitle); err != nil {
			log.Fatal("Failed to update todo:", err)
		}
		fmt.Println("Todo updated successfully")

	case "remove":
		removeCmd.Parse(os.Args[2:])
		if *removeId < 1 {
			log.Fatal("Invalid ID - must be positive integer")
		}
		if err := store.RemoveTodo(*removeId); err != nil {
			log.Fatal("Failed to remove todo:", err)
		}
		fmt.Println("Todo removed successfully")
	case "clean":
		cleanCmd.Parse(os.Args[2:])
		if err := store.Clean(); err != nil {
			log.Fatal("Failed to clean completed todos:", err)
		}
		fmt.Println("Completed todos removed successfully")

	case "list":
		listCmd.Parse(os.Args[2:])
		todos, err := store.GetTodos()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Todos:")
		for _, todo := range todos {
			status := " "
			if todo.Completed {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s\n",
				status,
				todo.Id,
				todo.Title,
			)
		}
	default:
		fmt.Printf("Unknown command: %s\n", opCommand)
		os.Exit(1)
	}
}
