import { API_ORIGIN } from "../constants";

export async function getAll() {
    return await fetch(`${API_ORIGIN}/todos`).then(res => res.json()).catch(err => {
        console.error("Error fetching todos:", err);
        throw err;
    });
}

export async function getByID(id: number) {
    await fetch(`${API_ORIGIN}/todos/${id}`).then(res => res.json()).catch(err => {
        console.error(`Error fetching todo with id ${id}:`, err);
        throw err;
    });
}


export async function create(title: string, description: string = "") {
    const newTodo = {
        title,
        description
    };

    return await fetch(`${API_ORIGIN}/todos`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(newTodo)
    }).then((res) => res.json()).catch(err => {
        console.error("Error creating todo:", err);
        throw err;
    });
}

export async function update(id: string | number, title: string, description: string, completed: boolean) {
    await fetch(`${API_ORIGIN}/todos/${id}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ title, description, completed })
    }).then(res => res.json()).catch(err => {
        console.error(`Error updating todo with id ${id}:`, err);
        throw err;
    });
}

export async function updateCompleted(id: string | number, completed: boolean) {
    await fetch(`${API_ORIGIN}/todos/${id}`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ completed })
    }).then(res => res.json()).catch(err => {
        console.error(`Error updating todo with id ${id}:`, err);
        throw err;
    });
}

export async function deleteByID(id: number) {
    await fetch(`${API_ORIGIN}/todos/${id}`, {
        method: "DELETE"
    }).catch(err => {
        console.error(`Error deleting todo with id ${id}:`, err);
        throw err;
    });
}

export async function deleteAll() {
    await fetch(`${API_ORIGIN}/todos`, {
        method: "DELETE"
    }).catch(err => {
        console.error(`Error deleting all todos`, err);
        throw err;
    });
}

