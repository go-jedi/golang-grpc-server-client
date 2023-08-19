package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/grpc-new/pb"
	"github.com/rob-bender/grpc-new/pkg/logging"
	"google.golang.org/grpc/codes"
)

type AuthPostgres struct {
	db     *sqlx.DB
	logger logging.Logger
}

func NewAuthPostgres(db *sqlx.DB, logger logging.Logger) *AuthPostgres {
	return &AuthPostgres{
		db:     db,
		logger: logger,
	}
}

func (r *AuthPostgres) SignUp(authDto *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	var uid string
	var isCreateUser bool
	authDtoJson, err := json.Marshal(authDto)
	if err != nil {
		r.logger.Errorf("ошибка конвертации authDto, %s", err)
		return &pb.SignUpResponse{Status: int32(codes.Internal), Message: "ошибка конвертации authDto", Result: false}, fmt.Errorf("ошибка конвертации userForm, %s", err)
	}
	err = r.db.QueryRow("SELECT uid($1)", 8).Scan(&uid)
	if err != nil {
		r.logger.Errorf("ошибка выполнения функции uid из базы данных, %s", err)
		return &pb.SignUpResponse{Status: int32(codes.Internal), Message: "ошибка выполнения функции uid из базы данных", Result: false}, fmt.Errorf("ошибка выполнения функции uid из базы данных, %s", err)
	}
	err = r.db.QueryRow("SELECT user_create($1, $2)", authDtoJson, uid).Scan(&isCreateUser)
	if err != nil {
		r.logger.Errorf("ошибка выполнения функции user_create из базы данных, %s", err)
		return &pb.SignUpResponse{Status: int32(codes.Internal), Message: "ошибка выполнения функции user_create из базы данных", Result: false}, fmt.Errorf("ошибка выполнения функции user_create из базы данных, %s", err)
	}
	r.logger.Info("успешное создание пользователя")
	return &pb.SignUpResponse{Status: int32(codes.OK), Message: "Успешное создание пользователя", Result: isCreateUser}, nil
}

func (r *AuthPostgres) GetUser(username string, password string) ([]*pb.User, error) {
	var user []*pb.User
	var userByte []byte
	err := r.db.QueryRow("SELECT user_get_data($1, $2)", username, password).Scan(&userByte)
	if err != nil {
		r.logger.Errorf("ошибка выполнения функции user_get_data из базы данных, %s", err)
		return []*pb.User{}, fmt.Errorf("ошибка выполнения функции user_get_data из базы данных, %s", err)
	}
	err = json.Unmarshal(userByte, &user)
	if err != nil {
		r.logger.Errorf("ошибка конвертации в функции GetUser, %s", err)
		return []*pb.User{}, fmt.Errorf("ошибка конвертации в функции GetUser, %s", err)
	}
	r.logger.Info("успешное получение пользователя")
	return user, nil
}

func (r *AuthPostgres) AddRefreshToken(id int, refreshToken string, expiresAt time.Time) (int, error) {
	_, err := r.db.Exec("SELECT refresh_token_add($1, $2, $3)", id, refreshToken, expiresAt)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции refresh_token_add из базы данных, %s", err)
	}
	return http.StatusOK, nil
}
