package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/grpc-new/pb"
	"github.com/rob-bender/grpc-new/pkg/logging"
)

type TodoAuth interface {
	SignUp(authDto *pb.SignUpRequest) (*pb.SignUpResponse, error)
	GetUser(username string, password string) ([]*pb.User, error)
	AddRefreshToken(id int, refreshToken string, expiresAt time.Time) (int, error)
}

type TodoUser interface {
	GetUserProfile(userDto *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error)
}

type Repository struct {
	TodoAuth
	TodoUser
}

func NewRepository(db *sqlx.DB, logger logging.Logger) *Repository {
	return &Repository{
		TodoAuth: NewAuthPostgres(db, logger),
		TodoUser: NewUserPostgres(db, logger),
	}
}
