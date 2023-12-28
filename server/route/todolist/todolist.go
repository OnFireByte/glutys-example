package todolist

import (
	"errors"
	"server/reqcontext"
)

type Todo struct {
	Id        int
	Title     string
	Completed bool
}

var todos = map[string][]Todo{}

var acc = 0

func Add(user reqcontext.Username, title string) Todo {
	acc++
	todo := Todo{acc, title, false}
	todos[string(user)] = append(todos[string(user)], todo)
	return todo
}

func Get(user reqcontext.Username, id int) Todo {
	for _, todo := range todos[string(user)] {
		if todo.Id == id {
			return todo
		}
	}
	return Todo{}
}

func GetAll(user reqcontext.Username) []Todo {
	return todos[string(user)]
}

func Update(user reqcontext.Username, id int, title string, completed bool) (Todo, error) {
	for i, todo := range todos[string(user)] {
		if todo.Id == id {
			todos[string(user)][i] = Todo{id, title, completed}
			return todos[string(user)][i], nil
		}
	}
	return Todo{}, errors.New("todo not found")
}
