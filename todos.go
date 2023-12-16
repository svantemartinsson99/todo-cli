package main

import (
	"strconv"
	"log"
	"fmt"
)

type Todo struct {
	Id int
	Text string
	Priority int
	Tag string
}

func listTodos(todos []*Todo) {

	fmt.Printf("%-3s|%-10s|%-8s|%-s\n", "ID", "TAG", "PRIORITY", "TEXT")
	for _, t := range todos {
	fmt.Printf("%-3d|%-10s|%-8d|%-s\n", t.Id, t.Tag, t.Priority, t.Text)
	}
}

func addTodo(todos []*Todo, args map[string]string) []*Todo {
	var t Todo
	t.Id = generateId(todos)

	text, ok := args["text"]
	if !ok {
		log.Fatal("Could not add new todo, missing text parameter!")
	}
	t.Text = text
	
	t.Priority = 1
	priority, ok := args["priority"]
	if ok {
		priority, err := strconv.Atoi(priority)
		if err != nil {
			t.Priority = 1
			log.Print("Could not set priority, error: ", err, ". Setting priority to 1!")
		} else {
			t.Priority = priority
		}
	}

	tag, ok := args["tag"]
	if ok {
		t.Tag= tag
	}
	return append(todos, &t)
}

func removeTodo(todos []*Todo, args map[string]string) []*Todo {
	idStr, ok := args["id"]
	if !ok {
		log.Fatal("Could not remove todo, missing required parameter id!")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal("Provided id is not a valid integer!")
	}

	for i, t := range todos {
		if t.Id == id {
			if i != len(todos) - 1 {
				todos[i] = todos[len(todos)-1]
			}
			todos = todos[:len(todos)-1]
		}
	}
	return todos
}

func generateId(todos []*Todo) int {
	id := 1
	for i, t := range todos {
		if t.Id >= i {
			id = t.Id + 1
		}
	}
	return id
}

