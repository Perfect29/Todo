package service

import (
	"errors"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type ToDoList []Todo

type TodoService struct {
	toDos ToDoList
}

type Todo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

func (s *TodoService) RemoveTodo(id string) error {
	// var toDosCopy ToDoList
	// toDosCopy = append(toDosCopy, toDos[:id]...)
	// toDosCopy = append(toDosCopy, toDos[id + 1:]...)
	// toDos = toDosCopy
	index := -1
	for ind, todo := range s.toDos {
		if todo.ID == id {
			index = ind
			break
		}
	}
	if index == -1 {
		return errors.New("ID was not found")
	}

	var toDosCopy ToDoList
	toDosCopy = append(toDosCopy, s.toDos[:index]...)
	toDosCopy = append(toDosCopy, s.toDos[index+1:]...)
	s.toDos = toDosCopy
	return nil
}

func (s *TodoService) AddTodo(todo *Todo) error {
	// todo generate id, math.Rand
	todo.ID = uuid.NewString()
	s.toDos = append(s.toDos, *todo)
	return nil
}

func (s *TodoService) ShowListTodo() (ToDoList, error) {
	return s.toDos, nil
}
