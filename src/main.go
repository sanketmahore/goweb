package main

import (
	"crud/src/bookingimpl/rest"
	"fmt"
	"crud/src/auth_server"
)

func main() {

	fmt.Println("Crud operations....")

	go grpcserver.RunGRPC()

	rest.HandleRequests()

	
}
