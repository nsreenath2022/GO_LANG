package main
import ("fmt")


/*write a program that declares two variables
● one variable to store a VALUE of TYPE int8
○ assign to it the largest number possible, then print it
● one variable to store a VALUE of TYPE uint8
○ assign to it the largest number possible, then print it*/
func main () {
	var one int8 = 127
	var two uint8 = 255

	fmt.Println(one, two)
}