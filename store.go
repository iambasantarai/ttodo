package main

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Id          int64
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Store struct {
	conn *sql.DB
}

func (s *Store) Init() error {
	var err error
	s.conn, err = sql.Open("sqlite3", "./todos.db")
	if err != nil {
		return err
	}

	createTableStmt := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER NOT NULL PRIMARY KEY,
        title TEXT NOT NULL,
        completed INTEGER NOT NULL DEFAULT 0,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        completed_at TIMESTAMP
    );
    `
	_, err = s.conn.Exec(createTableStmt)
	return err
}

func (s *Store) GetTodos() ([]Todo, error) {
	rows, err := s.conn.Query(`
        SELECT id, title, completed, created_at, completed_at 
        FROM todos
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.CompletedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *Store) AddTodo(title string) error {
	insertStmt := `
    INSERT INTO todos(
        title, completed, created_at
    ) 
    VALUES(?, 0, ?);
    `
	_, err := s.conn.Exec(insertStmt, title, time.Now())
	return err
}

func (s *Store) RemoveTodo(id int64) error {
	result, err := s.conn.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("todo not found")
	}

	return nil
}

func (s *Store) ToggleTodo(id int64) error {
	tx, err := s.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var completed bool
	err = tx.QueryRow("SELECT completed FROM todos WHERE id = ?", id).Scan(&completed)
	if err != nil {
		return err
	}

	var toggleStmt string
	if !completed {
		toggleStmt = `UPDATE todos SET completed = 1, completed_at = ? WHERE id = ?`
		_, err = tx.Exec(toggleStmt, time.Now(), id)
	} else {
		toggleStmt = `UPDATE todos SET completed = 0, completed_at = NULL WHERE id = ?`
		_, err = tx.Exec(toggleStmt, id)
	}
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Store) UpdateTodo(id int64, title string) error {
	result, err := s.conn.Exec("UPDATE todos SET title = ? WHERE id = ?", title, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("todo not found")
	}

	return nil
}

func (s *Store) Clean() error {
	result, err := s.conn.Exec(
		"DELETE FROM todos WHERE completed = ? AND completed_at <= ?",
		1,
		time.Now(),
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("todos not found")
	}

	return nil
}

func (s *Store) Close() error {
	return s.conn.Close()
}
