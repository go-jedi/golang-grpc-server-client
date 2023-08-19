package service

import (
	"context"
	"fmt"
	"time"

	"github.com/rob-bender/grpc-new/internal/repository"
	"github.com/rob-bender/grpc-new/pb"
	"github.com/rob-bender/grpc-new/pkg/hash"
	"github.com/rob-bender/grpc-new/pkg/jwt"
	"github.com/rob-bender/grpc-new/pkg/logging"
	"google.golang.org/grpc/codes"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	repo       repository.TodoAuth
	jwtManager *jwt.JWTManager
	logger     logging.Logger
}

func NewAuthServer(r repository.TodoAuth, jwtManager *jwt.JWTManager, logger logging.Logger) *AuthServer {
	return &AuthServer{
		repo:       r,
		jwtManager: jwtManager,
		logger:     logger,
	}
}

func (s *AuthServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	s.logger.Info("запрос на создание пользователя")
	resGeneratePasswordHash, err := hash.GeneratePasswordHash(req.Password)
	if err != nil {
		s.logger.Errorf("ошибка выполнения функции GeneratePasswordHash, %s", err)
		return &pb.SignUpResponse{Status: int32(codes.Internal), Message: "ошибка выполнения функции GeneratePasswordHash", Result: false}, fmt.Errorf("ошибка выполнения функции GeneratePasswordHash, %s", err)
	}
	req.Password = resGeneratePasswordHash
	return s.repo.SignUp(req)
}

func (s *AuthServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	s.logger.Info("запрос на вход в аккаунт")
	resGeneratePasswordHash, err := hash.GeneratePasswordHash(req.Password)
	if err != nil {
		s.logger.Errorf("ошибка выполнения функции GeneratePasswordHash, %s", err)
		return &pb.SignInResponse{Status: int32(codes.Internal), Message: "ошибка выполнения функции GeneratePasswordHash", AccessToken: "", RefreshToken: ""}, fmt.Errorf("ошибка выполнения функции GeneratePasswordHash, %s", err)
	}

	user, err := s.repo.GetUser(req.Username, resGeneratePasswordHash)
	if err != nil {
		s.logger.Errorf("ошибка выполнения функции GetUser, %s", err)
		return &pb.SignInResponse{Status: int32(codes.Internal), Message: "ошибка выполнения функции GetUser", AccessToken: "", RefreshToken: ""}, fmt.Errorf("ошибка выполнения функции GetUser, %s", err)
	}
	if len(user) == 0 {
		return &pb.SignInResponse{Status: int32(codes.Internal), Message: "неправильный логин или пароль", AccessToken: "", RefreshToken: ""}, fmt.Errorf("неправильный логин или пароль")
	}

	accessToken, refreshToken, err := s.jwtManager.GenerateTokens(int(user[0].Id))
	if err != nil {
		s.logger.Errorf("ошибка выполнения функции GenerateTokens, %s", err)
		return &pb.SignInResponse{Status: int32(codes.Internal), Message: "ошибка выполнения функции GenerateTokens", AccessToken: "", RefreshToken: ""}, fmt.Errorf("ошибка выполнения функции GenerateTokens, %s", err)
	}

	_, err = s.repo.AddRefreshToken(int(user[0].Id), refreshToken, time.Now().Add(time.Hour*24*30))
	if err != nil {
		s.logger.Errorf("ошибка выполнения функции AddRefreshToken, %s", err)
		return &pb.SignInResponse{Status: int32(codes.Internal), Message: "ошибка выполнения функции AddRefreshToken", AccessToken: "", RefreshToken: ""}, fmt.Errorf("ошибка выполнения функции AddRefreshToken, %s", err)
	}

	return &pb.SignInResponse{
		Status:       int32(codes.OK),
		Message:      "Успешная авторизация",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
