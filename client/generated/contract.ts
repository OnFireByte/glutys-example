export type TodolistTodo = {
	Id: number
	Title: string
	Completed: boolean
	Tag: string[]
}

export type TodolistUpdateTodo = {
	Id: number
	Title: string
	Completed: boolean
	Tag: string[]
}

export type GlutysContract = {
	"math.fib": (n: number) => Promise<number>;
	"todolist.add": (title: string, tag: string[]) => Promise<TodolistTodo>;
	"todolist.bulkAdd": (titles: string[]) => Promise<TodolistTodo[]>;
	"todolist.get": (id: number) => Promise<TodolistTodo>;
	"todolist.getAll": () => Promise<TodolistTodo[]>;
	"todolist.update": (todo: (TodolistUpdateTodo | null)) => Promise<TodolistTodo>;
}
