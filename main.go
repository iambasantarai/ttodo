package main

import "fmt"

func main() {
	todos := Todos{}

	todos.add("Refactor code")
	todos.add("Learn react")
	todos.add("Write a unit test for s3explorer")
	fmt.Println("TODO: before")
	for _, todo := range todos {
		fmt.Printf("[%t] %s\n", todo.Completed, todo.Title)
	}

	fmt.Println("TODO: after")
	todos.delete(1)
	for _, todo := range todos {
		fmt.Printf("[%t] %s\n", todo.Completed, todo.Title)
	}
}
