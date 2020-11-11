package main

import (
	"testing"
	"time"
)

var all, er = NewAllTasks()

func TestAddTask(t *testing.T) {
	task := Task{"Test Title", "Test Body", time.Now(), false, 1}
	_, e := all.AddTask(&task)
	if e != nil {
		t.Error("Error! Can't be added")
	}
}

func TestDeleteTask(t *testing.T) {
	bef, _ := all.GetallTasks()
	var id int = 1
	e := all.DeleteTask(id)
	af, _ := all.GetallTasks()
	if e != nil && len(bef)-len(af) != 1 {
		t.Error("Error! Can't be delete")
	}
}

func TestChangeTask(t *testing.T) {
	var id int = 2
	nt := Task{"Changed Test Title", "Changed Task Body", time.Now(), true, id}
	_, err := all.ChangeTask(id, &nt)
	if err != nil {
		t.Error("Error! Can't be changed")
	}
}

func TestGetAllTasks(t *testing.T) {
	_, er := all.GetallTasks()
	if er != nil {
		t.Error("Error! Can't be got All the tasks")
	}
}
