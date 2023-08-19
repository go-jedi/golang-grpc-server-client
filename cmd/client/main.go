package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rob-bender/grpc-new/client"
	"github.com/rob-bender/grpc-new/pb"
	"github.com/rob-bender/grpc-new/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()
	log.Printf("dial server %s, TLS = %t", *serverAddress, *enableTLS)

	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	if *enableTLS {
		tlsCredentials, err := tls.LoadTLSCredentialsClient()
		if err != nil {
			log.Fatalf("cannot load TLS credentials: %v", err)
		}
		transportOption = grpc.WithTransportCredentials(tlsCredentials)
	}

	cc1, err := grpc.Dial(*serverAddress, transportOption)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	authClient := client.NewAuthClient(cc1)

	cc2, err := grpc.Dial(
		*serverAddress,
		transportOption,
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	userClient := client.NewUserClient(cc2)

	signUpRequest := &pb.SignInRequest{
		Username: "rob-bender1",
		Password: "24972497Vlad",
	}
	resSignIn, err := authClient.SignIn(signUpRequest)
	if err != nil {
		log.Fatalf("error in SignUp: %v", err)
	}
	fmt.Println("res -->", resSignIn)

	resGetUserProfile, err := userClient.GetUserProfile(&pb.GetUserProfileRequest{
		Id: 1,
	})
	if err != nil {
		log.Fatalf("error in GetUserProfile: %v", err)
	}
	fmt.Println("resGetUserProfile -->", resGetUserProfile)
}
