package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
)

const (
	c0 = iota
	c1 = iota
	c2 = iota
)

const (
	c3 = iota
	c4
	c5
	c6
	c7
	c8
)

func print_iota() {
	fmt.Println(c0, c1, c2)

	fmt.Println(c3, c4, c5)

	fmt.Println(c6, c7, c8)
}

func main() {
	fmt.Println("Welcome to assignment session:", cmplx.Sqrt(9))
	fmt.Println(rand.Intn(2))
	print_iota();
}