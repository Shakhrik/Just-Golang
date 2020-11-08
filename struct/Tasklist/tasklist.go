package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Task struct
type Task struct {
	title, body string
	createdAt   time.Time
	done        bool
	id          int
}

// Constructor for Task
func NewTask() Task {
	t := Task{}
	return t
}

type TaskManager interface {
	getDetails()
	getAllTasks()
	addTask()
	iscompleted()
}

// Tasks of Task
type Tasks struct {
	tasks []Task
}

func (t *Tasks) getAllTasks() []Task {
	return t.tasks
}
func (ts *Tasks) addTask(t *Task) {
	ts.tasks = append(ts.tasks, *t)
}
func (ts *Tasks) deleteTask(id int) {
	ts.tasks = append(ts.tasks[:id-1], ts.tasks[id:]...)
}
func (ts *Tasks) changeTask(id int) {
	update := &ts.tasks[id-1]
	if update.id == id {
		getDetails(update, ts)
	} else {
		fmt.Println("Can't be changed!")
	}

}
func (ts *Tasks) iscompleted(t *Task) {
	var answer string
	fmt.Print("Is this task is completed?:[Y/N]")
	fmt.Scan(&answer)
	if answer == "Y" || answer == "y" {
		t.done = true
	} else if answer == "N" || answer == "n" {
		t.done = false
	} else {
		fmt.Println("Please enter the mentioned chars!")
	}
}
func getDetails(todo *Task, ts *Tasks) {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Enter todo title:")
	title, _ := input.ReadString('\n')
	fmt.Print("Enter todo body:")
	body, _ := input.ReadString('\n')
	ts.iscompleted(todo)
	todo.title = title
	todo.createdAt = time.Now()
	todo.body = body
	todo.id = len(ts.tasks) + 1
}
func menu() {
	fmt.Println("-----Task list----")
	fmt.Println("1. Add task")
	fmt.Println("2.Update existing task")
	fmt.Println("3.Delete existing task")
	fmt.Println("4.Show all tasks")
	fmt.Println("0.Exit")
}

func main() {
	task := &Task{}
	tasks := Tasks{}
	var input int = 1
	inputt := bufio.NewReader(os.Stdin)
	for input != 0 {
		menu()
		fmt.Scan(&input)
		switch input {
		case 1:
			getDetails(task, &tasks)
			tasks.addTask(task)
			break
		case 2:
			fmt.Print("Please enter the title of task you want to update:")
			title, _ := inputt.ReadString('\n')
			var doesntexist bool = true
			d := tasks.getAllTasks()
			for _, v := range d {
				if v.title == title {
					tasks.changeTask(v.id)
					fmt.Println("Successfully changed!")
					doesntexist = false
					break
				}
			}
			if doesntexist {
				fmt.Println("The task with this title doesn't exist!!!")
			}
			break
		case 3:
			var doesntexist bool = true
			fmt.Println("Enter Task title:")
			title, _ := inputt.ReadString('\n')
			d := tasks.getAllTasks()
			for _, v := range d {
				fmt.Println("v.title", v.title)
				fmt.Println("entered title", title)
				if title == v.title {
					tasks.deleteTask(v.id)
					fmt.Println("DELETED successfully")
					doesntexist = false
					break
				}
			}
			if doesntexist {
				fmt.Println("Task with this title doesn't exist")
			}
			break
		case 4:
			alltasks := tasks.getAllTasks()
			if len(alltasks) == 0 {
				fmt.Println("You have no any contact")
			}
			for i, v := range alltasks {
				fmt.Print(i+1, ".Task:")
				fmt.Println("\nTitle:", v.title, "Body:", v.body, "Created at:", v.createdAt.Format("2006-01-02 15:04"),
					"\nCompleted:", v.done)
			}
			break
		}
	}
}
