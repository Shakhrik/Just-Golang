package main

import (
	"testing"
	"time"
)

var all = NewAllTasks()

func TestAddTask(t *testing.T) {
	var task = Task{"Greeting", "Hello hello", time.Now(), false, 1}
	var before int = len(all.tasks)
	all.AddTask(&task)
	var after int = len(all.tasks)
	if after-before != 1 {
		t.Error("Error! Can't be added")
	}
}
func TestDeleteTask(t *testing.T) {
	all.tasks[0] = Task{"Laptop", "Hello again", time.Now(), true, 1}
	var before int = len(all.tasks)
	all.DeleteTask(1)
	var after int = len(all.tasks)
	if before-after != 1 {
		t.Error("Error! Can't be deleted")
	}
}

func TestChangeTask(t *testing.T) {
	task := Task{"Laptop", "Hello again", time.Now(), true, 1}
	all.tasks = append(all.tasks, task)
	all.ChangeTask(1)
	if task == all.tasks[0] {
		t.Error("Error! Can't be changed")
	}

}
func TestGetAllTasks(t *testing.T) {
	if len(all.tasks) != len(all.getAllTasks()) {
		t.Error("Error! Getalltasks function is not workin properly")
	}
}
