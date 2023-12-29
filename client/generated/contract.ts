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
	"todolist.bulkAdd": (arg0: (string | null)[]) => Promise<TodolistTodo[]>;
	"todolist.get": (arg0: number) => Promise<TodolistTodo>;
	"todolist.getAll": () => Promise<TodolistTodo[]>;
	"todolist.update": (arg0: (TodolistUpdateTodo | null)) => Promise<TodolistTodo>;
	"math.fib": (arg0: number) => Promise<number>;
	"todolist.add": (arg0: string, arg1: string[]) => Promise<TodolistTodo>;
}
