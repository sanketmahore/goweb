package main

import (
	"context"
	"crud/src/authentication"
	"crud/src/bookingimpl/dao"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	authentication.UnimplementedLoginServiceServer
}

func (*server) Authenticate(ctx context.Context, input *authentication.Credentials) (*authentication.Result, error) {
	fmt.Println("In grpcServer : Authenticate start")
	d := dao.NewBookingDao()
	res, err := d.AuthenticateUser(input)
	fmt.Println("In grpcServer : Authenticate end")
	return &authentication.Result{
		Outcome: res,
	}, err
}

func main() {
	fmt.Println("Authentication server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
		return
	}

	s := grpc.NewServer()

	authentication.RegisterLoginServiceServer(s, &server{})

	err1 := s.Serve(lis)
	if err1 != nil {
		log.Fatalf("Failed to serve %v", err)
		return
	}
}
