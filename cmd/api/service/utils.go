package service

import (
	_ "github.com/lib/pq"
)

type ToDoList []Todo

type Todo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          int    `json:"id"`
}
