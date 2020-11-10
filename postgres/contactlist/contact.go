package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Contact struct {
	firstName, lastName, phoneNumber string
}
type Functions interface {
	AddContact()
	UpdateContact()
	DeleteContact()
	getAllcontacts()
}
type AllContacts struct {
	db *sql.DB
}

func NewAllContacts() (*AllContacts, error) {
	c := AllContacts{}
	var err error
	c.db, err = sql.Open("postgres", "user=postgres port=5432 dbname=postgres password=1 host=localhost sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &c, nil
}
func (c *AllContacts) AddContact(nc *Contact) (*Contact, error) {
	sqlStatement := `INSERT INTO contact(firstname, lastname, phonenumber) VALUES($1, $2, $3)`
	_, err := c.db.Exec(sqlStatement, nc.firstName, nc.lastName, nc.phoneNumber)
	if err != nil {
		return nil, err
	}
	return nc, nil
}

func (c *AllContacts) UpdateContact(nc *Contact, id int) *Contact {
	sqlStatement := `UPDATE contact SET firstname=$1, lastname=$2, phonenumber=$3 WHERE id=$4`
	_, err := c.db.Exec(sqlStatement, nc.firstName, nc.lastName, nc.phoneNumber, id)
	if err != nil {
		panic(err)
	}
	return nc
}
func (c *AllContacts) DeleteContact(id int) error {
	sqlDelete := `DELETE FROM contact WHERE id=$1`
	_, err := c.db.Exec(sqlDelete, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *AllContacts) GetAllcontacts() ([]Contact, error) {
	row, err := c.db.Query("SELECT * FROM contact")
	if err != nil {
		panic(err)
		return nil, err
	}
	var f, l, p string
	var id int
	var contacts []Contact
	for row.Next() {
		err := row.Scan(&id, &f, &l, &p)
		if err != nil {
			panic(err)
			return nil, err
		}
		contacts = append(contacts, Contact{f, l, p})
	}
	return contacts, nil
}

func main() {
	connStr := "user=postgres port=5432 dbname=postgres password=1 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the database")
	// c := Contact{"Elyor", "Ismatov", "951248795"}
	cs, err := NewAllContacts()
	// _, err = cs.AddContact(&c)
	// newf := cs.UpdateContact(&c, 2)
	// newf := cs.DeleteContact(16)
	_, err = cs.GetAllcontacts()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
