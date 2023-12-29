package todolist

import (
	"errors"

	"server/reqcontext"
)

type Todo struct {
	Id        int
	Title     string
	Completed bool
	Tag       []string
}

var todos = map[string][]Todo{}

var acc = 0

func Add(user reqcontext.Username, title string, tag []string) Todo {
	acc++
	// NOTE: You must return empty slice instead of nil slice, since nil slice will be encoded as null in JSON. This is the encode/json's behavior.
	todo := Todo{acc, title, false, []string{}}
	todos[string(user)] = append(todos[string(user)], todo)
	return todo
}

func BulkAdd(user reqcontext.Username, titles []string) ([]Todo, error) {
	var res []Todo
	for _, title := range titles {
		res = append(res, Add(user, title, nil))
	}
	return res, nil
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

func Update(user reqcontext.Username, todo *Todo) (Todo, error) {
	// pointer type is act as nullable
	// type *Todo is converted to Todo | null 
	if todo == nil {
		return Todo{}, errors.New("todo is null")
	}
	for i, v := range todos[string(user)] {
		if v.Id == todo.Id {
			todos[string(user)][i] = *todo
			return todos[string(user)][i], nil
		}
	}
	return Todo{}, errors.New("todo not found")
}
