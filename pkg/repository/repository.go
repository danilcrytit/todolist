package repository

import (
	"github.com/jmoiron/sqlx"
	"todolist"
)

type Authorization interface {
	CreateUser(user todolist.User) (int, error)
	GetUser(username, password string) (todolist.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
