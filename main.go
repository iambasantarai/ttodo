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

	cliArgs = cliArgs[2:]

	switch opCommand {
	case "add":
		title := flag.String("t", "Untitled", "title for todo")
		description := flag.String("d", "", "description for todo")

		flag.Parse()

		fmt.Println(*title)
		fmt.Println(*description)
	case "toggle":
		id := flag.Int("i", 0, "id of todo")

		flag.Parse()

		fmt.Println(*id)
	case "update":
		id := flag.Int("i", 0, "id of todo")
		title := flag.String("t", "Untitled", "title for todo")
		descriptionPtr := flag.String("d", "", "description for todo")

		flag.Parse()

		fmt.Println(*id)
		fmt.Println(*title)
		fmt.Println(*descriptionPtr)
	case "remove":
		id := flag.Int("i", 0, "id of todo")

		flag.Parse()

		fmt.Println(*id)
	default:
		fmt.Printf("Unknown command: %s", cliArgs[1])
		os.Exit(1)
	}
}
