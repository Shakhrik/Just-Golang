package main

import (
	"context"
	"log"
	"testing"

	"github.com/Shakhrik/Just-Golang/contactlist_grpc/proto"
	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc"
)

func TestAdd(t *testing.T) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewContactManagerClient(conn)
	newContact := &proto.Contact{
		Id:          1,
		FirstName:   "Shakhriyor",
		LastName:    "Ismatov",
		PhoneNumber: "+998933559909",
	}
	res, e := client.Add(context.Background(), newContact)
	if e != nil {
		panic("Error!Can't Added")
	}
	log.Println("Response:", res)

}
func TestUpdate(t *testing.T) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewContactManagerClient(conn)
	updContact := &proto.Contact{
		Id:          1,
		FirstName:   "Updated Name",
		LastName:    "Updated Lastname",
		PhoneNumber: "+99899999999",
	}
	res, e := client.Update(context.Background(), updContact)
	if e != nil {
		panic("Error! Can't Updated")
	}
	log.Println("Res:", res)
}
func TestDelete(t *testing.T) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewContactManagerClient(conn)
	delCon := &proto.ContactId{Id: 1}
	res, e := client.Delete(context.Background(), delCon)
	if e != nil {
		panic("Error!Can't deleted!")
	}
	log.Println(res)
}
func TestGetAllContacts(t *testing.T) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewContactManagerClient(conn)
	res, e := client.GetAllContacts(context.Background(), &empty.Empty{})
	if e != nil {
		panic("Error!Can't get all contacts...")
	}
	log.Println("Res:", res)
}
