package user

import (
	"context"

	pb "github.com/Soumik43/grpc-user-service/api/user"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	repo Repository
}

func NewUserServiceServer(repo Repository) *UserServiceServer {
	return &UserServiceServer{repo: repo}
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	return s.repo.GetUser(req.GetId())
}

func (s *UserServiceServer) GetUsers(ctx context.Context, req *pb.UserIdList) (*pb.Users, error) {
	return s.repo.GetUsers(req.GetIds())
}
