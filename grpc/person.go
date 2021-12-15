package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc_test/genproto"
	"grpc_test/model"
)

type PersonService struct {
	grpc.UnimplementedPersonServiceServer
}

func (s *PersonService) GetPersonInfo(context.Context, *grpc.PersonName) (*grpc.Person, error) {
	var person *model.Person
	person = &model.Person{
		ID:    1,
		Name:  "John",
		Email: "example@email.com",
		Phones: []*model.PhoneNumber{
			{
				Number: "13542687562",
				Type:   model.HOME,
			},
			{
				Number: "13764852159",
				Type:   model.MOBILE,
			},
		},
		LastUpdated: timestamppb.Now(),
	}
	var personInfo *grpc.Person
	personInfo = person.ToGRPCPerson()
	return personInfo, nil
}
