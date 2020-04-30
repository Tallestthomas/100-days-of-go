package main

import "fmt"

// In order to use Iota, we need the constants to be at a package level
const (
	// each time 'iota' is used, it increments
	first  = iota
	second = iota
	// you can use iota for constant expressions
)

// Each group of "iota" declaration resets
const (
	// this would be 7
	third = iota + 6
	// this would be 4 (2 bitshifted 4 times)
	fourth = 2 << iota
)

func main() {
	// will print 0, 1
	fmt.Println(first, second)
	fmt.Println(third, fourth)
}
