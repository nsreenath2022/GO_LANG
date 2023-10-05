package main

import (
	"fmt"
)

func main() {
	name , age, fee := "Sreenath Natarajan", 42, 89.65

	fmt.Printf("\n%s of type %T", name, name)
	fmt.Printf("\n%d of type %T", age, age)
	fmt.Printf("\n%f of type %T", fee, fee)
}