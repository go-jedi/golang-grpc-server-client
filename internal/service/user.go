package service

import (
	"context"

	"github.com/rob-bender/grpc-new/internal/repository"
	"github.com/rob-bender/grpc-new/pb"
	"github.com/rob-bender/grpc-new/pkg/logging"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo   repository.TodoUser
	logger logging.Logger
}

func NewUserServer(r repository.TodoUser, logger logging.Logger) *UserServer {
	return &UserServer{
		repo:   r,
		logger: logger,
	}
}

func (s *UserServer) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	s.logger.Info("запрос на получение профиля пользователя")
	return s.repo.GetUserProfile(req)
}
