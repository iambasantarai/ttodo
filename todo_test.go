package main

import (
	"testing"
)

func TestTodoOperations(t *testing.T) {
	t.Run("add todo", func(t *testing.T) {
		todos := &Todos{}
		title := "Test Todo"
		description := "This is a test todo."

		todos.add(title, description)

		if len(*todos) != 1 {
			t.Errorf("Expected 1 todo, got %d", len(*todos))
		}

		if (*todos)[0].Title != title {
			t.Errorf("Expected title %s, got %s", title, (*todos)[0].Title)
		}

		if (*todos)[0].Description != description {
			t.Errorf("Expected description %s, got %s", description, (*todos)[0].Description)
		}
	})

	t.Run("delete todo", func(t *testing.T) {
		todos := &Todos{
			{Title: "Todo 1", Description: "Description 1", Completed: false},
			{Title: "Todo 2", Description: "Description 2", Completed: false},
		}

		err := todos.remove(0)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if len(*todos) != 1 {
			t.Errorf("Expected 1 todo, got %d", len(*todos))
		}

		if (*todos)[0].Title != "Todo 2" {
			t.Errorf("Expected title Todo 2, got %s", (*todos)[0].Title)
		}
	})

	t.Run("toggle todo", func(t *testing.T) {
		todos := &Todos{
			{Title: "Todo 1", Description: "Description 1", Completed: false},
		}

		err := todos.toggle(0)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !(*todos)[0].Completed {
			t.Errorf("Expected todo to be completed")
		}

		if (*todos)[0].CompletedAt == nil {
			t.Errorf("Expected CompletedAt to be set")
		}
	})

	t.Run("update todo", func(t *testing.T) {
		todos := &Todos{
			{Title: "Todo 1", Description: "Description 1", Completed: false},
		}

		title := "Updated Todo"
		description := "Updated Description"
		err := todos.update(0, title, description)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if (*todos)[0].Title != title {
			t.Errorf("Expected title %s, got %s", title, (*todos)[0].Title)
		}

		if (*todos)[0].Description != description {
			t.Errorf("Expected description %s, got %s", description, (*todos)[0].Description)
		}
	})
}

func TestValidateIndex(t *testing.T) {
	todos := &Todos{
		{Title: "Todo 1", Description: "Description 1", Completed: false},
	}

	testCases := []struct {
		index       int
		expectError bool
		description string
	}{
		{0, false, "Valid index"},
		{1, true, "Index out of range"},
		{-1, true, "Negative index"},
	}

	for _, tc := range testCases {
		err := todos.validateIndex(tc.index)
		if (err != nil) != tc.expectError {
			t.Errorf("%s: Expected error: %v, got: %v", tc.description, tc.expectError, err != nil)
		}
	}
}
