package main

import (
	"os"
	"log"
	"strings"
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
	  	default:
			  log.Fatal("Invalid action!")
	}

	saveTodos(todos)
}

func parseCliInput() (string, map[string]string) {
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
