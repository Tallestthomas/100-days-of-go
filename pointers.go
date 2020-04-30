package main

import "fmt"

func main() {
	// Initialize a variable
	firstName := "Arthur"
	fmt.Println(firstName)

	// Create and AddressOf pointer
	ptr := &firstName
	// Print the pointer's memory address and dereferenced value
	fmt.Println(ptr, *ptr)

	// Change the variables value
	firstName = "Tricia"
	// The pointer's memory addres will remain the same, but the dereferenced value will have changed
	fmt.Println(ptr, *ptr)
}
