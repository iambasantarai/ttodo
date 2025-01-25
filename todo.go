package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title, description string) {
	todo := Todo{
		Title:       title,
		Description: description,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("Invalid index")
	}

	return nil
}

func (todos *Todos) remove(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) update(index int, title, description string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	t[index].Description = description

	return nil
}

func (todos *Todos) list() {
	for idx, todo := range *todos {
		fmt.Printf("%d. [%t] %s\n %s\n\n", idx, todo.Completed, todo.Title, todo.Description)
	}
}
