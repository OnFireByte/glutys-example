package todolist

type Todo struct {
	Id        int
	Title     string
	Completed bool
}

var todos = []Todo{}

var acc = 0

func Add(title string) Todo {
	acc++
	todo := Todo{acc, title, false}
	todos = append(todos, todo)
	return todo
}

func Get(id int) Todo {
	for _, todo := range todos {
		if todo.Id == id {
			return todo
		}
	}
	return Todo{}
}

func GetAll() []Todo {
	return todos
}

func Update(id int, title string, completed bool) Todo {
	for i, todo := range todos {
		if todo.Id == id {
			todos[i] = Todo{id, title, completed}
			return todos[i]
		}
	}
	return Todo{}
}
