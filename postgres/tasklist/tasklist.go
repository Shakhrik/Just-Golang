package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Task struct
type Task struct {
	title, body string
	createdAt   time.Time
	done        bool
	id          int
}
type Tasks struct {
	db *sql.DB
}

func NewTask() Task {
	t := Task{}
	return t
}
func NewAllTasks() (*Tasks, error) {
	ts := Tasks{}
	var err error
	ts.db, err = sql.Open("postgres", "user=postgres port=5432 dbname=postgres password=1 host=localhost sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &ts, nil
}

type TaskManager interface {
	GetallTasks()
	AddTask()
	DeleteTask()
	ChangeTask()
}

func (t *Tasks) GetallTasks() ([]Task, error) {
	row, er := t.db.Query("SELECT * FROM task")
	if er != nil {
		return nil, er
	}
	var id int
	var tit, b string
	var d bool
	var ct time.Time
	ts := []Task{}
	for row.Next() {
		err := row.Scan(&id, &tit, &b, &d, &ct)
		if err != nil {
			return nil, err
		}
		ts = append(ts, Task{tit, b, ct, d, id})
	}
	return ts, nil
}
func (ts *Tasks) AddTask(t *Task) (*Task, error) {
	sqlAdd := `INSERT INTO task(id, title, body, done, created_at) VALUES($1, $2, $3, $4, $5)`
	_, er := ts.db.Exec(sqlAdd, t.id, t.title, t.body, t.done, t.createdAt)
	if er != nil {
		return nil, er
	}
	log.Println("Added successfully")
	return t, nil
}
func (ts *Tasks) ChangeTask(id int, nt *Task) (*Task, error) {
	sqlUpdate := `UPDATE task SET id=$1, title=$2, body=$3, done=$4 WHERE id=$5`
	_, err := ts.db.Exec(sqlUpdate, id, nt.title, nt.body, nt.done, id)
	if err != nil {
		return nil, err
	}
	return nt, nil
}

func (ts *Tasks) DeleteTask(id int) error {
	sqlDelete := `DELETE FROM task WHERE id=$1`
	_, err := ts.db.Exec(sqlDelete, id)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// connStr := "user=postgres port=5432 dbname=postgres password=1 host=localhost sslmode=disable"
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Successfully connected to the database")
	// t := Task{id: 3, title: "Changed title", body: "Changed body", done: false, createdAt: time.Now()}
	// ts, _ := NewAllTasks()
	// tfdas := &Tasks{}
	// r, errrr := ts.GetallTasks()
	// fmt.Println(r, errrr)

	// if er != nil {
	// 	panic(er)
	// }
	// _, e := ts.AddTask(&t)
	// _, e := ts.ChangeTask(3, &t)
	// ts.DeleteTask(0)

}
