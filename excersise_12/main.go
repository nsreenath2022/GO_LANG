package main
import (
	"fmt"
)
/*
write a program that uses print verbs to show the following numbers
● 747
● 911
● 90210
as
● decimal
● binary
● hexadecimal
*/
func main () {
	a , b , c := 747, 911, 90210

	fmt.Printf("%d, %b, %#X ", a, a, a)
	fmt.Printf("\n%d, %b, %#X ", b, b, b)
	fmt.Printf("\n%d, %b, %#X ", c, c, c)
}