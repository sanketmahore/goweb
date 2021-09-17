package main

import (
	"fmt"
	"crud/src/bookingimpl"
)

func main() {

	fmt.Println("Crud operations....")

	bookingimpl.HandleRequests()
}