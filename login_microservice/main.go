package main

import (
	"github.com/sanketmahore/goweb_microservices/login_microservice/rest"
	"fmt"
)

func main() {

	fmt.Println("Crud operations....")

	// go grpcserver.RunGRPC()

	rest.HandleRequests()

}
