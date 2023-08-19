package client

import (
	"context"
	"fmt"
	"time"

	"github.com/rob-bender/grpc-new/pb"
	"google.golang.org/grpc"
)

type UserClient struct {
	service pb.UserServiceClient
}

func NewUserClient(cc *grpc.ClientConn) *UserClient {
	service := pb.NewUserServiceClient(cc)
	return &UserClient{service}
}

func (client *UserClient) GetUserProfile(body *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("111")
	res, err := client.service.GetUserProfile(ctx, body)
	if err != nil {
		return res, err
	}
	return res, nil
}
