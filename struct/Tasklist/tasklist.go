package main

import (
	"bufio"
	"fmt"
	"os"
)

// Tasklist struct
type Tasklist struct {
	nameOftodo, id string
}

func getDetails(todo *Tasklist) {

	input := bufio.NewReader(os.Stdin)
	fmt.Print("Enter todo name:")
	title, _ := input.ReadString('\n')
	todo.nameOftodo = title
}
func main() { 
	justTodo := &Tasklist{}
	getDetails(justTodo)
}
