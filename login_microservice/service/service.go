package service

import (
	"goweb_microservices/login_microservice/authentication"
	"context"
	"goweb_microservices/login_microservice/domain"
	"fmt"
	"strings"

	"google.golang.org/grpc"
)

//var bookingDao domain.BookingDao

type service struct {
}

func NewLoginService() domain.LoginService {
	return &service{}
}

func (s *service) Login(creds *domain.Credentials) (bool, error) {
	conn, err := grpc.Dial("172.16.238.11:50051", grpc.WithInsecure())
	if err != nil {
		return false, fmt.Errorf("failed to connect %v", err)
	}

	defer conn.Close()

	c := authentication.NewLoginServiceClient(conn)
	req := &authentication.Credentials{
		Username: creds.Username,
		Password: creds.Password,
	}

	res, err := c.Authenticate(context.Background(), req)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, fmt.Errorf("error in RPC %v", err)
	}

	return res.Outcome, nil
}
