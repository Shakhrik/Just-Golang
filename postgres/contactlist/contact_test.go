package main

import "testing"

var all, err = NewAllContacts()

func TestAddConctact(t *testing.T) {
	var c = &Contact{"TestUser", "Testing", "2154212054"}
	res, e := all.AddContact(c)
	if e != nil || res.firstName != c.firstName {
		t.Error("Error! Can't be add", e)
	}
}
func TestUpdateContact(t *testing.T) {
	newCon := &Contact{"NewTestUser", "NewTestLastname", "933358859"}
	_, e := all.UpdateContact(newCon, 23)
	if e != nil {
		t.Error("Error!Can't be updated")
	}
}
func TestDeleteContact(t *testing.T) {
	id := 20
	befo, _ := all.GetAllcontacts()
	e := all.DeleteContact(id)
	afte, _ := all.GetAllcontacts()
	if e != nil && len(befo)-len(afte) != 1 {
		t.Error("Error! Can't be deleted")
	}
}
func TestGetAllContacst(t *testing.T) {
	_, er := all.GetAllcontacts()
	if er != nil {
		t.Error("Error! Can't be get all contacts")
	}
}
