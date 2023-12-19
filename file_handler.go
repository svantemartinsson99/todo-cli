package main

import (
  "os"
  "encoding/json"
  "log"
)

var todoFile string
const FILE_FORMAT = ".dat"
const DEFAULT_TODOS_FILE_NAME = "todos"

func saveTodos(todos []*Todo) {
  json, err := json.Marshal(todos)
  if err != nil {
    log.Fatal("Could not convert todos to JSON! CHANGES ARE NOT SAVED!")
  }

  f, err := os.Create("./" + todoFile)
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
  if todoFile == "" {
    todoFile = DEFAULT_TODOS_FILE_NAME + FILE_FORMAT 
  }

  data, e := os.ReadFile("./" + todoFile)
  if e != nil {
    if os.IsNotExist(e) {
      log.Print("No file found with name " + todoFile + ", a new file will be created.")	
      var todos []*Todo
      return todos
    } else {
      log.Fatal("Could not read todos file, error: ", e)
    }
  }
  var todos []*Todo
  err := json.Unmarshal(data, &todos)
  if err != nil {
    log.Fatal("Could not parse json, error: ", e)
  }
  return todos
}
