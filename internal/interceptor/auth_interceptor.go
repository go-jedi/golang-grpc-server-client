package interceptor

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/rob-bender/grpc-new/pkg/helpers"
	"github.com/rob-bender/grpc-new/pkg/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
	userCtx             = "userId"
)

var withoutUserIdentity []string = []string{
	"/auth_service.AuthService/SignUp",
	"/auth_service.AuthService/SignIn",
}

type AuthInterceptor struct {
	jwtManager *jwt.JWTManager
}

func NewAuthInterceptor(jwtManager *jwt.JWTManager) *AuthInterceptor {
	return &AuthInterceptor{jwtManager}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		resContainsString := helpers.ContainsString(withoutUserIdentity, info.FullMethod)
		if !resContainsString {
			err := interceptor.userIdentity(ctx, info.FullMethod)
			fmt.Println("err -->", err)
			if err != nil {
				return nil, err
			}
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)

		resContainsString := helpers.ContainsString(withoutUserIdentity, info.FullMethod)
		if !resContainsString {
			err := interceptor.userIdentity(stream.Context(), info.FullMethod)
			if err != nil {
				return err
			}
		}

		return handler(srv, stream)
	}
}

func (interceptor *AuthInterceptor) userIdentity(ctx context.Context, method string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "метаданные не предоставляются")
	}
	if len(md[authorizationHeader]) == 0 {
		return status.Errorf(codes.Unauthenticated, "пустой заголовок")
	}

	headerParts := strings.Split(md[authorizationHeader][0], " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return status.Errorf(codes.Unauthenticated, "неверный заголовок авторизации")
	}
	if len(headerParts[1]) == 0 {
		return status.Errorf(codes.Unauthenticated, "токен пуст")
	}

	userId, err := interceptor.jwtManager.ParseToken(headerParts[1])
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}
	if userId < 1 {
		return status.Errorf(codes.Unauthenticated, "идентификатор пользователя имеет недопустимый тип")
	}

	return nil
}
