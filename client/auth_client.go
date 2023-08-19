package client

import (
	"context"
	"time"

	"github.com/rob-bender/grpc-new/pb"
	"google.golang.org/grpc"
)

type AuthClient struct {
	service pb.AuthServiceClient
}

func NewAuthClient(cc *grpc.ClientConn) *AuthClient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{service}
}

func (client *AuthClient) SignUp(body *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := client.service.SignUp(ctx, body)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (client *AuthClient) SignIn(body *pb.SignInRequest) (*pb.SignInResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := client.service.SignIn(ctx, body)
	if err != nil {
		return res, err
	}
	return res, nil
}
