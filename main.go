package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	cliArgs := os.Args
	if len(cliArgs) < 2 {
		fmt.Println("expected 'add', 'update', 'remove', or 'toggle' subcommand")
		os.Exit(1)
	}

	opCommand := cliArgs[1]

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCmd.String("t", "Untitled", "title for todo")
	addDescription := addCmd.String("d", "", "description for todo (optional)")

	toggleCmd := flag.NewFlagSet("toggle", flag.ExitOnError)
	toggleId := toggleCmd.Int("i", 0, "id of todo to toggle status")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCmd.Int("i", 0, "id of todo to update")
	updateTitle := updateCmd.String("t", "Untitled", "new title for todo")
	updateDescription := updateCmd.String("d", "", "new description for todo (optional)")

	removeCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	removeId := removeCmd.Int("i", 0, "id of todo to toggle status")

	cliArgs = cliArgs[2:]

	switch opCommand {
	case "add":
		addCmd.Parse(os.Args[2:])

		fmt.Println(*addTitle)
		fmt.Println(*addDescription)
	case "toggle":
		toggleCmd.Parse(os.Args[2:])

		fmt.Println(*toggleId)
	case "update":
		updateCmd.Parse(os.Args[2:])

		fmt.Println(*updateId)
		fmt.Println(*updateTitle)
		fmt.Println(*updateDescription)
	case "remove":
		removeCmd.Parse(os.Args[2:])

		fmt.Println(*removeId)
	default:
		fmt.Printf("Unknown command: %s", cliArgs[1])
		os.Exit(1)
	}
}
