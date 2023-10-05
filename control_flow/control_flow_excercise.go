package main

import (
	"fmt"
	"math/rand"
)

func excercise_one(){
	x := rand.Intn(250)
	fmt.Println(" x =", x)

	if x >=0 && x< 101{
		fmt.Println("between 0 and 100")
	} else if x >= 101 && x <= 200{
		fmt.Println("between 101 and 200")
	} else if x > 200 &&  x < 251 {
		fmt.Println("between 201 and 250")
	} else {
		fmt.Println("This is more then 250")
	}

	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
}


func excercise_two(){
	
	x := rand.Intn(250)
	fmt.Println(" x =", x)

	
	switch {
	case x >=0 && x<= 100 :
			fmt.Println("between 0 and 100")
	case  x >= 101 && x <= 200:
			fmt.Println("between 101 and 200")
	case x > 200 &&  x < 251 :
			fmt.Println("between 201 and 250")
	default:
			fmt.Println("This is more then 250")
	}
}

func excercise_three(){
	fmt.Println("This is dummy function ")
	
}

func init(){
	fmt.Println("This is init function from main.go")
}


func excercise_four(){
	x := rand.Intn(10)
	y:= rand.Intn(10)

	fmt.Printf(" x = %v\t y = %v\n", x, y)

	if x <4 && y < 4{
		fmt.Println("x and y both less than 4")
	} else if x > 6 && y > 6{
		fmt.Println("x and y both greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x  greater than or equal to 4 and less than or equal to 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("none of the previous statements are met")
	}
	
}


func excercise_five(){
	x := rand.Intn(10)
	y:= rand.Intn(10)

	fmt.Printf(" x = %v\t y = %v\n", x, y)

	switch {
	case x <4 && y < 4:
			fmt.Println("x and y both less than 4")
	case x > 6 && y > 6:
			fmt.Println("x and y both greater than 6")
	case x >= 4 && x <= 6:
			fmt.Println("x  greater than or equal to 4 and less than or equal to 6")
	case y != 5:
			fmt.Println("y is not 5")
	default:
			fmt.Println("none of the previous statements are met")
	}
	
}

func excercise_six () {
	
}

func main () {
	//excercise_one()
	//excercise_two()
	//excercise_three()
	//excercise_four()
	//excercise_five()
	excercise_six()
}