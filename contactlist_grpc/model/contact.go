package model

import (
	"github.com/jmoiron/sqlx"

	cont "github.com/Shakhrik/Just-Golang/contactlist_grpc/proto"

	_ "github.com/lib/pq"
)

type ContactManager interface {
	AddContact(c *cont.Contact) error
	UpdateContact(c *cont.Contact, id int64) error
	DeleteContact(id int64)
	GetAllcontacts() ([]*cont.Contact, error)
}
type AllContacts struct {
	db *sqlx.DB
}

func NewAllContacts() (*AllContacts, error) {
	c := AllContacts{}
	var err error
	c.db, err = sqlx.Connect("postgres", "user=postgres port=5432 dbname=postgres password=1 host=localhost sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *AllContacts) AddContact(nc *cont.Contact) error {
	sqlStatement := `INSERT INTO contact(firstname, lastname, phonenumber) VALUES($1, $2, $3)`
	c.db.MustExec(sqlStatement, nc.FirstName, nc.LastName, nc.PhoneNumber)
	return nil
}

func (c *AllContacts) UpdateContact(nc *cont.Contact, id int64) error {
	sqlStatement := `UPDATE contact SET firstname=$1, lastname=$2, phonenumber=$3 WHERE id=$4`
	c.db.MustExec(sqlStatement, nc.FirstName, nc.LastName, nc.PhoneNumber, id)
	return nil
}

func (c *AllContacts) DeleteContact(id int64) error {
	sqlDelete := `DELETE FROM contact WHERE id=$1`
	c.db.MustExec(sqlDelete, id)
	return nil
}

func (c *AllContacts) GetAllcontacts() ([]*cont.Contact, error) {
	sqlstatement := "SELECT * FROM contact"
	var contacts []*cont.Contact
	err := c.db.Select(&contacts, sqlstatement)
	return contacts, err
}
