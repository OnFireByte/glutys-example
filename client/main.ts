import axios, { AxiosError } from "axios";
import { CreateAPIClient } from "glutys-client";
import { GlutysContract, TodolistTodo } from "./generated/contract";

const instance = axios.create({
    baseURL: "http://localhost:8080/api",
    headers: {
        username: "john",
    },
});

const api = CreateAPIClient<GlutysContract>(instance);

async function main() {
    try {
        console.log("Fetching");

        const req: Promise<TodolistTodo>[] = [];

        for (let i = 0; i < 10; i++) {
            const p = api.todolist.add(`Todo ${i}`);
            req.push(p);
        }

        await Promise.all(req);

        const test = await api.todolist.getAll();
        console.log(test);
    } catch (e) {
        if (e instanceof AxiosError) {
            console.log(e.message);
            console.log(e.response!.data);
        }
    }
}

main();
