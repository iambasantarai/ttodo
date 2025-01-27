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

	cliArgs := os.Args
	if len(cliArgs) < 2 {
		fmt.Println("expected 'add', 'update', 'remove', or 'toggle' subcommand")
		os.Exit(1)
	}

	opCommand := cliArgs[1]

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCmd.String("t", "Untitled", "title for todo")

	toggleCmd := flag.NewFlagSet("toggle", flag.ExitOnError)
	toggleId := toggleCmd.Int64("i", -1, "id of todo to toggle status")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCmd.Int64("i", -1, "id of todo to update")
	updateTitle := updateCmd.String("t", "Untitled", "new title for todo")

	removeCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	removeId := removeCmd.Int64("i", -1, "id of todo to toggle status")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	cliArgs = cliArgs[2:]

	switch opCommand {
	case "add":
		addCmd.Parse(os.Args[2:])

		store.AddTodo(*addTitle)
	case "toggle":
		toggleCmd.Parse(os.Args[2:])

		store.ToggleTodo(*toggleId)
	case "update":
		updateCmd.Parse(os.Args[2:])

		store.UpdateTodo(*updateId, *updateTitle)
	case "remove":
		removeCmd.Parse(os.Args[2:])

		store.RemoveTodo(*removeId)
	case "list":
		listCmd.Parse(os.Args[2:])

		store.GetTodos()
	default:
		fmt.Printf("Unknown command: %s", cliArgs[1])
		os.Exit(1)
	}
}
