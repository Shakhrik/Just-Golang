package main

import (
	"net"

	"context"

	"github.com/Shakhrik/Just-Golang/contactlist_grpc/model"
	proto "github.com/Shakhrik/Just-Golang/contactlist_grpc/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	cs model.ContactManager
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	ts, _ := model.NewAllContacts()
	srv := grpc.NewServer()
	proto.RegisterContactManagerServer(srv, &server{ts})
	reflection.Register(srv)
	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) Add(ctx *context.Context, contact *proto.Contact) (*proto.Contact, error) {
	err := s.cs.AddContact(contact)
	if err != nil {
		return nil, err
	}
	return contact, nil
}
func (s *server) Update(ctx *context.Context, nc *proto.Contact) (*proto.Status, error) {
	err := s.cs.UpdateContact(nc, nc.GetId())
	if err != nil {
		return nil, err
	}
	return &proto.Status{Status: "updated successfully"}, nil
}
func (s *server) Delete(ctx *context.Context, id *proto.ContactId) (*proto.Status, error) {
	s.cs.DeleteContact(id.GetId())
	return &proto.Status{Status: "deleted successfully"}, nil
}
func (s *server) GetAllContacts(ctx context.Context, e *empty.Empty) (*proto.Contacts, error) {
	contacts, err := s.cs.GetAllcontacts()
	if err != nil {
		panic(err)
		return nil, err
	}
	return &proto.Contacts{Contacts: contacts}, nil
}
