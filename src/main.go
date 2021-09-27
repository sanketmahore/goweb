package main

import (
	"crud/src/bookingimpl/rest"
	"fmt"
)

func main() {

	fmt.Println("Crud operations....")

	rest.HandleRequests()
}
