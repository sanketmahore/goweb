package main

import (
	"fmt"
	"goweb_microservices/admin_microservice/src/rest"
)

func main() {

	fmt.Println("Admin Microservice...")

	// go grpcserver.RunGRPC()

	rest.HandleRequests()

}
