package main

import (
  "os"
  "log"
  "strings"
  "fmt"
)

func main() {
  action, args := parseCliInput()
	
  todos := loadTodos()

  switch {
    case action == "add" || action == "a":
      todos = addTodo(todos, args)
    case action == "remove" ||action == "r":
      todos = removeTodo(todos, args)
    case action == "list" || action == "l":
      listTodos(todos, args)
    case action == "start" || action == "s":
      setStatus(todos, args, "In progress")
    case action == "done" || action == "d":
      setStatus(todos, args, "Done")
    case action == "reset" || action == "rs":
      setStatus(todos, args, "Not started")
    case action == "update" || action == "u":
      updateTodo(todos, args)
    case action == "help" || action == "h":
      printHelp()
    default:
      log.Fatal("Invalid action!")
  }

  saveTodos(todos)
}

func printHelp() {
  fmt.Printf("%-12s| %-s\n\n", "Command", "Description")
  fmt.Printf("%-12s| %-s\n", "add or a", "Adds a new todo. Arguments: text, priority (optional), tag (optional)")
  fmt.Printf("%-12s| %-s\n", "remove or r", "Removes a todo. Argument: id")
  fmt.Printf("%-12s| %-s\n", "list or l", "Lists todos. Can also filter with arguments. Arguments: tag (optional), status (optional)")
  fmt.Printf("%-12s| %-s\n", "start or s", "Sets status to ''In progress''. No arguments.")
  fmt.Printf("%-12s| %-s\n", "done or d", "Sets status to ''Done''. No arguments.")
  fmt.Printf("%-12s| %-s\n", "reset or rs", "Sets status to ''Not started''. No arguments.")
}

func parseCliInput() (string, map[string]string) {

  if len(os.Args) < 2 {
    action := "l"
    return action, make(map[string]string, 0)
  }
  action := os.Args[1]

  cliArgs := os.Args[2:]
  invalidArgMsg := " is skipped since it is not a valid argument format! Argument have the format arg=value!"
  args := make(map[string]string)

  for _, arg := range cliArgs {
    if !strings.Contains(arg, "=")  {
      log.Print(arg, invalidArgMsg)
      continue
    }
    argArr := strings.Split(arg, "=")
    if len(argArr) != 2 {
      log.Print(arg, invalidArgMsg)
      continue;
    }
    if argArr[0] == "file" || argArr[0] == "f" {
      todoFile = argArr[1] + FILE_FORMAT
      continue
    }
    args[argArr[0]] = argArr[1]
  }
  return action, args
}
