package main

import (
	"fmt"
	"crud/src/bookingimpl"
)

func main() {

	fmt.Println("Curd operations....")

	bookingimpl.HandleRequests()
}
