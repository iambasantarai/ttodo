package main

import "fmt"

func main() {
	todos := Todos{}

	todos.add("Refactor code")
	todos.add("Learn react")

	for _, todo := range todos {
		fmt.Printf("[ ] %s\n", todo.Title)
	}
}
