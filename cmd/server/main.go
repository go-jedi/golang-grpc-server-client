package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rob-bender/grpc-new/internal/interceptor"
	"github.com/rob-bender/grpc-new/internal/repository"
	"github.com/rob-bender/grpc-new/internal/service"
	"github.com/rob-bender/grpc-new/pb"
	"github.com/rob-bender/grpc-new/pkg/jwt"
	"github.com/rob-bender/grpc-new/pkg/postgres"

	"github.com/rob-bender/grpc-new/pkg/logging"
	"github.com/rob-bender/grpc-new/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	signingKey = "pH9MBWu01iJk6ZNQTrHiRUdCLhTy4xykA1zHnTBULa8h0"
	// tokenTTL   = 15 * time.Minute <- prod
	tokenTTL = 1 * time.Hour // <- for dev
)

func runGRPCServer(
	authServer pb.AuthServiceServer,
	userServer pb.UserServiceServer,
	jwtManager *jwt.JWTManager,
	enableTLS bool,
	listener net.Listener,
	logger logging.Logger,
) error {
	authInterceptor := interceptor.NewAuthInterceptor(jwtManager)
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(authInterceptor.Unary()),
		grpc.StreamInterceptor(authInterceptor.Stream()),
	}

	if enableTLS {
		logger.Info("loadTLSCredentials initializing")
		tlsCredentials, err := tls.LoadTLSCredentialsServer()
		if err != nil {
			logger.Fatalf("cannot load TLS credentials: %s", err.Error())
		}
		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	}

	grpcServer := grpc.NewServer(serverOptions...)

	pb.RegisterAuthServiceServer(grpcServer, authServer)
	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	logger.Infof("Start GRPC server at %s, TLS = %t", listener.Addr().String(), enableTLS)
	return grpcServer.Serve(listener)
}

func main() {
	port := flag.Int("port", 0, "the server port")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()

	logging.Init()
	logger := logging.GetLogger()
	logger.Info("logger initialized")
	logger.Info("config initializing")
	initConfig(logger)
	logger.Info("db initializing")
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "24972497Vlad",
		DBName:   "test_db",
		SslMode:  "disable",
	})
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}
	logger.Info("repository initializing")
	repos := repository.NewRepository(db, logger)

	jwtManager := jwt.NewJWTManager(signingKey, tokenTTL)
	authServer := service.NewAuthServer(repos, jwtManager, logger)
	userServer := service.NewUserServer(repos, logger)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatalf("Failed to start server %v", err)
	}

	err = runGRPCServer(authServer, userServer, jwtManager, *enableTLS, listener, logger)
	if err != nil {
		logger.Fatal("cannot start server in runGRPCServer: ", err)
	}
}

func initConfig(logger logging.Logger) {
	if err := godotenv.Load(); err != nil {
		logger.Fatal(err)
	}
	for _, k := range []string{"DATABASE_HOST", "DATABASE_PORT", "DATABASE_USERNAME", "DATABASE_PASSWORD", "DATABASE_NAME", "DATABASE_SSL_MODE"} {
		if _, ok := os.LookupEnv(k); !ok {
			logger.Fatalf("set environment variable -> %s", k)
		}
	}
}
