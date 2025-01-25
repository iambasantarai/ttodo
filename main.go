package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	todos := Todos{}
	store := NewStore[Todos]("todos.json")
	store.load(&todos)

	cliArgs := os.Args
	if len(cliArgs) < 2 {
		fmt.Println("expected 'add', 'update', 'remove', or 'toggle' subcommand")
		os.Exit(1)
	}

	opCommand := cliArgs[1]

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCmd.String("t", "Untitled", "title for todo")

	toggleCmd := flag.NewFlagSet("toggle", flag.ExitOnError)
	toggleId := toggleCmd.Int("i", -1, "id of todo to toggle status")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCmd.Int("i", -1, "id of todo to update")
	updateTitle := updateCmd.String("t", "Untitled", "new title for todo")

	removeCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	removeId := removeCmd.Int("i", -1, "id of todo to toggle status")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	cliArgs = cliArgs[2:]

	switch opCommand {
	case "add":
		addCmd.Parse(os.Args[2:])

		todos.add(*addTitle)
	case "toggle":
		toggleCmd.Parse(os.Args[2:])

		todos.toggle(*toggleId)
	case "update":
		updateCmd.Parse(os.Args[2:])

		todos.update(*updateId, *updateTitle)
	case "remove":
		removeCmd.Parse(os.Args[2:])

		todos.remove(*removeId)
	case "list":
		listCmd.Parse(os.Args[2:])

		todos.list()
	default:
		fmt.Printf("Unknown command: %s", cliArgs[1])
		os.Exit(1)
	}

	store.save(todos)
}
