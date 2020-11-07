package main

import (
	"fmt"
)

// Contact struct
type Contact struct {
	firstName, lastName, phoneNumber string
	id                               int
}

// Functions interface
type Functions interface {
	addContact()
	updateContact()
	deleteContact()
	getAllcontacts()
}

func getDetails(contact *Contact, a AllContacts) {
	fmt.Print("Firstname: ")
	fmt.Scan(&contact.firstName)
	fmt.Print("Lastname: ")
	fmt.Scan(&contact.lastName)
	fmt.Print("Phone number: ")
	fmt.Scan(&contact.phoneNumber)
	contact.id = len(a.contacts) + 1

}
func menu() {
	fmt.Println("-----Contact list----")
	fmt.Println("1. Add new contact")
	fmt.Println("2.Update existing contact")
	fmt.Println("3.Delete existing contact")
	fmt.Println("4.Show all contacts")
	fmt.Println("0.Exit")

}

// AllContacts struct
type AllContacts struct {
	contacts []Contact
}

// NewAllContacts constructor
func NewAllContacts() AllContacts {
	c := AllContacts{}
	c.contacts = []Contact{}
	return c
}
func (c *AllContacts) addContact(newContact *Contact) {
	c.contacts = append(c.contacts, *newContact)

}
func (c *AllContacts) updateContact(newCon *Contact) *Contact {
	oldCont := c.contacts[newCon.id]
	oldCont.firstName = newCon.firstName
	oldCont.lastName = newCon.lastName
	oldCont.phoneNumber = newCon.phoneNumber
	return &oldCont
}
func (c *AllContacts) getAllcontacts() []Contact {
	return c.contacts
}
func (c *AllContacts) deleteContact(id int) {
	c.contacts = append(c.contacts[:id-1], c.contacts[id:]...)
}
func main() {
	var input int = 1
	c := &Contact{}
	p := AllContacts{}
	for input != 0 {
		menu()
		fmt.Scan(&input)
		switch input {
		case 1:
			getDetails(c, p)
			p.addContact(c)
			break
		case 3:
			var name string
			var doesntexist bool = true
			fmt.Println("Enter Contact name:")
			fmt.Scan(&name)
			d := p.getAllcontacts()
			for _, v := range d {
				if name == v.firstName {
					p.deleteContact(v.id)
					fmt.Println("DELETED successfully")
					doesntexist = false
					break
				}
			}
			if doesntexist {
				fmt.Println("Contact with this name doesn't exist")
			}
			break
		case 4:
			allConts := p.getAllcontacts()
			if len(allConts) == 0 {
				fmt.Println("You have no any contact")
			}
			for i, v := range allConts {
				fmt.Print(i+1, ".CONTACT:")
				fmt.Println("\nFirstname:", v.firstName, "\nLastname:", v.lastName, "\nPhone Number:", v.phoneNumber)
			}
			break
		}
	}
}
