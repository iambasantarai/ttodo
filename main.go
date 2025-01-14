package main

import "fmt"

func main() {
	todos := Todos{}

	todos.add("Refactor code", "My code definetly needs a proper refactoring")
	todos.add("Learn react", "Frontend gg")
	todos.add("Write a unit test for s3explorer", "")
	fmt.Println("TODO: before")
	for _, todo := range todos {
		fmt.Printf("[%t] %s\n %s\n\n", todo.Completed, todo.Title, todo.Description)
	}

	fmt.Println("TODO: after")
	todos.toggle(2)
	todos.delete(1)
	todos.edit(0, "Message her", "You should message her man")
	for _, todo := range todos {
		fmt.Printf("[%t] %s\n %s\n\n", todo.Completed, todo.Title, todo.Description)
	}
}
