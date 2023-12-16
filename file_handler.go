package main

import (
	"os"
	"encoding/json"
	"log"
)

const TODO_FILE = "todos.dat"

func saveTodos(todos []*Todo) {
	json, err := json.Marshal(todos)
	if err != nil {
		log.Fatal("Could not convert todos to JSON! CHANGES ARE NOT SAVED!")
	}

	f, err := os.Create("./" + TODO_FILE)
	if err != nil {
		panic(err)
	}
	defer f.Close()	
	_, err = f.Write(json)
	if err != nil {
		log.Print("Failed to save file, error: ", err)
	}
}

func loadTodos() []*Todo {
	data, e := os.ReadFile("./" + TODO_FILE)
	if e != nil {
		log.Fatal("Could not read todos file, error: ", e)
	}
	var todos []*Todo
	err := json.Unmarshal(data, &todos)
	if err != nil {
		log.Fatal("Could not parse json, error: ", e)
	}
	return todos
}