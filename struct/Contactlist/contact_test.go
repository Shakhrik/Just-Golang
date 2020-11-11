package main

import "testing"

var all = NewAllContacts()

func TestAddingNewContact(t *testing.T) {
	c := Contact{"Shakhriyor", "Ismatov", "998933559909", 1}
	saved := all.addContact(&c)
	if saved.id != c.id {
		t.Error("Error! Can't be added")
	}
}
func TestUpdateContact(t *testing.T) {
	all.contacts[0] = Contact{"Barello", "Tomas", "53443344", 0}
	s := Contact{"Jon", "Boris", "949595494", 0}
	update := all.updateContact(&s)
	if update == &s {
		t.Error("Error! Can't be updated", update, s)
	}

}
func TestGetAllContacts(t *testing.T) {
	if len(all.contacts) != len(all.getAllcontacts()) {
		t.Error("Error! Not same")
	}
}
func TestDeleteContact(t *testing.T) {
	all.contacts[0] = Contact{"Barello", "Tomas", "53443344", 0}
	all.deleteContact(1)
	if len(all.contacts) != 0 {
		t.Error("Error! Can't be deleted")
	}
}
