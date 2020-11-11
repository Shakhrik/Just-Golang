package main

import "testing"

var all, err = NewAllContacts()
var c = &Contact{"TestUser", "Testing", "43535345335"}

func TestAddConctact(t *testing.T) {
	res, e := all.AddContact(c)
	if e != nil || res.firstName != c.firstName {
		t.Error("Error! Can't be add", e)
	}
}
