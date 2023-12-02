export type TodolistTodo = {
	Id: number
	Title: string
	Completed: boolean
}

export type GlutysContract = {
	"todolist.get": (arg0: number) => Promise<TodolistTodo>;
	"todolist.getAll": () => Promise<TodolistTodo[]>;
	"todolist.update": (arg0: number, arg1: string, arg2: boolean) => Promise<TodolistTodo>;
	"math.fib": (arg0: number) => Promise<number>;
	"todolist.add": (arg0: string) => Promise<TodolistTodo>;
}
