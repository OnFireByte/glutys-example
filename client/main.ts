import { AxiosError } from "axios";
import { CreateAPIClient } from "glutys-client";
import { GlutysContract } from "./generated/contract";

const api = CreateAPIClient<GlutysContract>("http://localhost:8080/api");

async function main() {
    try {
        console.log("Fetching");

        for (let i = 0; i < 10; i++) {
            api.todolist.add(`Todo ${i}`);
        }

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
