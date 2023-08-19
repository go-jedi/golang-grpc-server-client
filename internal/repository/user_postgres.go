package repository

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/grpc-new/pb"
	"github.com/rob-bender/grpc-new/pkg/logging"
	"google.golang.org/grpc/codes"
)

type UserPostgres struct {
	db     *sqlx.DB
	logger logging.Logger
}

func NewUserPostgres(db *sqlx.DB, logger logging.Logger) *UserPostgres {
	return &UserPostgres{
		db:     db,
		logger: logger,
	}
}

func (r *UserPostgres) GetUserProfile(userDto *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	var userProfile []*pb.User
	var userProfileByte []byte
	err := r.db.QueryRow("SELECT user_get_profile($1)", userDto.Id).Scan(&userProfileByte)
	if err != nil {
		r.logger.Errorf("ошибка выполнения функции user_get_profile из базы данных, %s", err)
		return &pb.GetUserProfileResponse{Status: int32(codes.Internal), Message: "ошибка выполнения функции user_get_profile из базы данных", Result: []*pb.User{}}, fmt.Errorf("ошибка выполнения функции user_get_profile из базы данных, %s", err)
	}
	err = json.Unmarshal(userProfileByte, &userProfile)
	if err != nil {
		r.logger.Errorf("ошибка конвертации в функции GetUserProfile, %s", err)
		return &pb.GetUserProfileResponse{Status: int32(codes.Internal), Message: "ошибка конвертации в функции GetUserProfile", Result: []*pb.User{}}, fmt.Errorf("ошибка конвертации в функции GetUserProfile, %s", err)
	}
	r.logger.Info("успешное получение профиля пользователя")
	return &pb.GetUserProfileResponse{Status: int32(codes.OK), Message: "успешное получение профиля пользователя", Result: userProfile}, nil
}
