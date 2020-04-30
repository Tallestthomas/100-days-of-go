package main

import "fmt"

func main() {
	// Constants must be available during compilation, not run time
	// Implicit declaration
	const c = 3
	fmt.Println(c + 3)

	// Go dynamically evaluates a constant's type
	fmt.Println(c + 1.2)

	// Explicit declaration
	const b int = 3
}
