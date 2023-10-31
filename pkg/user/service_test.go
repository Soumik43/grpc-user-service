package user

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/Soumik43/grpc-user-service/api/user"
)

func setup() *UserServiceServer {
	repo := NewInMemoryUserRepository()
	return NewUserServiceServer(repo)
}

func TestGetUser(t *testing.T) {
	s := setup()
	user, err := s.GetUser(context.Background(), &pb.UserId{Id: 1})

	if err != nil {
		t.Errorf("Error is not nil: %v", err)
	}

	expected := &pb.User{
		Id:      1,
		Fname:   "Steve",
		City:    "LA",
		Phone:   1234567890,
		Height:  5.8,
		Married: true,
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("Expected %v, got %v", expected, user)
	}
}

func TestGetUserNonexistent(t *testing.T) {
	s := setup()
	_, err := s.GetUser(context.Background(), &pb.UserId{Id: 20})

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestGetUsers(t *testing.T) {
	s := setup()
	users, err := s.GetUsers(context.Background(), &pb.UserIdList{Ids: []int32{1, 2, 3}})

	if err != nil {
		t.Errorf("Error is not nil: %v", err)
	}

	expected := &pb.Users{
		Users: []*pb.User{
			{
				Id:      1,
				Fname:   "Steve",
				City:    "LA",
				Phone:   1234567890,
				Height:  5.8,
				Married: true,
			},
			{
				Id:      2,
				Fname:   "Shimmy",
				City:    "NYC",
				Phone:   9876543210,
				Height:  5.2,
				Married: false,
			},
			{
				Id:      3,
				Fname:   "Jane",
				City:    "HongKong",
				Phone:   1234509876,
				Height:  5.1,
				Married: true,
			},
		},
	}

	if !reflect.DeepEqual(users, expected) {
		t.Errorf("Expected %v, got %v", expected, users)
	}
}

func TestGetUsersSomeNonexistent(t *testing.T) {
	s := setup()
	users, err := s.GetUsers(context.Background(), &pb.UserIdList{Ids: []int32{1, 20}})

	if err != nil {
		t.Errorf("Error is not nil: %v", err)
	}

	expected := &pb.Users{
		Users: []*pb.User{
			{
				Id:      1,
				Fname:   "Steve",
				City:    "LA",
				Phone:   1234567890,
				Height:  5.8,
				Married: true,
			},
		},
	}

	if !reflect.DeepEqual(users, expected) {
		t.Errorf("Expected %v, got %v", expected, users)
	}
}
