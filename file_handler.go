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
  homeDir, e := os.UserHomeDir()
  if e != nil {
    log.Fatal("Could not get users home directory. Needed to save todos to file. Error: ", e)
  }

  f, err := os.Create(homeDir + "/.todo_cli/" + todoFile)
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

  homeDir, e := os.UserHomeDir()
  if e != nil {
    log.Fatal("Could not get users home directory. Needed to save todos to file. Error: ", e)
  }

  data, e := os.ReadFile(homeDir + "/.todo_cli/" + todoFile)
  if e != nil {
    if os.IsNotExist(e) {
      log.Print("No file found with name " + todoFile + ", a new file will be created.")	
      if e := os.MkdirAll(homeDir + "/.todo_cli", os.ModePerm); e != nil {
        log.Fatal("Could not create directory to save todos file! Error: ", e)
      }
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
