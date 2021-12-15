package model

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc_test/genproto"
)

type PhoneType int32

const (
	MOBILE PhoneType = 1
	HOME   PhoneType = 2
	WORK   PhoneType = 3
)

type PhoneNumber struct {
	Number string    `json:"number"`
	Type   PhoneType `json:"type"`
}
type Person struct {
	Name        string                 `json:"name"`
	ID          int32                  `json:"id"`
	Email       string                 `json:"email"`
	Phones      []*PhoneNumber         `json:"phones"`
	LastUpdated *timestamppb.Timestamp `json:"last_updated"`
}

func (person Person) ToGRPCPerson() *grpc.Person {
	p := new(grpc.Person)
	var phoneNumber []*grpc.Person_PhoneNumber
	for _, v := range person.Phones {
		phoneNumber = append(phoneNumber, v.ToGRPCPhoneNumber())
	}
	p.Id = person.ID
	p.Name = person.Name
	p.Email = person.Email
	p.LastUpdated = person.LastUpdated
	p.Phones = phoneNumber
	return p
}

func (p PhoneNumber) ToGRPCPhoneNumber() *grpc.Person_PhoneNumber {
	pn := new(grpc.Person_PhoneNumber)
	pn.Number = p.Number
	pn.Type = p.Type.ToGRPCPhoneType()
	return pn
}
func (pt PhoneType) ToGRPCPhoneType() grpc.Person_PhoneType {
	if pt == MOBILE {
		return grpc.Person_MOBILE
	} else if pt == HOME {
		return grpc.Person_HOME
	} else {
		return grpc.Person_WORK
	}
}
