package main

import (
  "strconv"
  "log"
  "fmt"
  "errors"
)

type Todo struct {
  Id int
  Text string
  Priority int
  Tag string
  Status string
}

func listTodos(todos []*Todo, args map[string]string) {
  l := filterTodos(todos, args)

  fmt.Printf("%-3s|%-10s|%-8s|%-12s|%-s\n", "ID", "TAG", "PRIORITY", "STATUS", "TEXT")
  for _, t := range l {
    fmt.Printf("%-3d|%-10s|%-8d|%-12s|%-s\n", t.Id, t.Tag, t.Priority, t.Status, t.Text)
  }
}

func filterTodos(todos []*Todo, args map[string]string)  []*Todo {
  var l []*Todo
  var tagFilter string
  var statusFilter string
  
  tag, ok := args["tag"]
  if ok {
    tagFilter = tag
  }

  status, ok := args["status"]
  if ok {
    statusFilter = status
  }

  for _, tPtr := range todos {
    if tagFilter != "" && tPtr.Tag != tagFilter {
      continue
    }
    if statusFilter != "" && tPtr.Status != statusFilter {
      continue
    }
    l = append(l, tPtr)
  } 
  return l
}

func validateStatus(status string) bool {
  return status == "Not started" || status == "In progress" || status == "Done"
}

func updateTodo(todos []*Todo, args map[string]string) {
  idStr, ok := args["id"]
  if !ok {
    log.Fatal("Missing required parameter id")
  }
  id, e := strconv.Atoi(idStr)
  if e != nil {
    log.Fatal("Invalid id, error: ", e)
  }

  todo, e := getTodo(id, todos)
  if e != nil {
    log.Fatal("Could not get todo, error: ", e)
  }

  text, ok := args["text"]
  if ok {
    todo.Text = text
  }

  tag, ok := args["tag"]
  if ok {
    todo.Tag = tag
  }

  prioStr, ok := args["priority"]
  if ok {
    prio, e := strconv.Atoi(prioStr)
    if e == nil {
      todo.Priority = prio
    } else {
      log.Print("Priority value is invalid, will not be updated. Error: ", e)
    }
  }
}

func getTodo(id int, todos []*Todo) (*Todo, error) {
  for _, t := range todos {
    if t.Id == id {
      return t, nil
    }
  }
  return nil, errors.New("Could not find any todo with id " + strconv.Itoa(id))
}

func setStatus(todos []*Todo, args map[string]string, status string) {
  if !validateStatus(status) {
    log.Fatal("Invalid status")
  }

  idStr, ok := args["id"]
  if !ok {
    log.Fatal("Missing required parameter id!")
  }

  id, e := strconv.Atoi(idStr)
  if e != nil {
    log.Fatal("Invalid id, error: ", e)
  }

  todo, e := getTodo(id, todos)
  if e != nil {
    log.Fatal("Could not find todo with id " + idStr + ". Error: ", e)
  }
  todo.Status = status
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
  t.Status = "Not started"
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

