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

func excercise_six(){
	count :=0
	for count < 100 {
		
		x := rand.Intn(10)
		y:= rand.Intn(10)

		fmt.Printf(" x = %v\t y = %v  count = %v \t", x, y, count)

		switch {
		case x < 4 && y < 4:
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
		count++
	}
}


func excercise_seven(){

	for i := 0; i <42; i++{
		x := rand.Intn(5)

		switch x{
		case 0:
			fmt.Println(" Case 0 - start or restart, counter = ", i)
		case 1:
			fmt.Println(" Case 1 - runnng process, counter = ", i)
		case 2:
			fmt.Println(" Case 2 - modifying process counter = ", i)
		case 3:
			fmt.Println(" Case 3 - printing process counter = ", i)
		case 4:
			fmt.Println(" Case 4 - stopping process counter = ", i)
		default:
			fmt.Println(" default - erroring process counter = ", i)

		}
	}
	
}


func main () {
	//excercise_one()
	//excercise_two()
	//excercise_three()
	//excercise_four()
	//excercise_five()
	//excercise_six()
	//excercise_seven()
	// excercise_eight()
	// excercise_nine()
	// excercise_ten()

	// excercise_nine_a()
	// excercise_nested_loop()
	// excercise_map()

	//excercise_comma_ok()

	//excercise_statement_statement()
	excercise_print_boolean_keys()
}

func excercise_print_boolean_keys(){
	fmt.Println("true && true ==>", true && true)
	fmt.Println("true && false ==>", true && false)
	fmt.Println("true || true ==>", true || true)
	fmt.Println("true || false ==>", true|| false)
	fmt.Println("!true ==>", !true)
}

func excercise_statement_statement(){
	cnt := 0
	for i := 0; i < 100; i++{
		if x := rand.Intn(5); x == 3{
			fmt.Println("x is 3")
			cnt++
		}
	}
	fmt.Println("count = ", cnt)
}

func excercise_comma_ok() {
	mp := map [string]int {
		"James": 42,
		"MoneyPenny": 32,
	}

	for key,val := range mp {
		fmt.Printf("name = %v and age = %v\n", key, val)
	}

	age := mp["James"]
	fmt.Println("James age = ", age)

	age1 := mp["Q"]
	fmt.Println("Q's age = ", age1)

	// comma ok is used to check if item present in map
	if val, ok := mp["James"]; ok{
		fmt.Println("James is present and here is val = ", val)
	} else {
		fmt.Println(" James is not present")
	}

	if _, check := mp["Q"]; check{
		fmt.Println("Q is present")
	} else {
		fmt.Println("Q is not present")
	}
}

func excercise_map(){
	mp := map [string]int {
		"James": 42,
		"MoneyPenny": 32,
	}

	for key,val := range mp {
		fmt.Printf("name = %v and age = %v\n", key, val)
	}

	james_age := mp["James"]
	fmt.Println("James age = ", james_age)
}

func excercise_nested_loop() {
	fmt.Println("------------------------------------")
	for i := 0; i < 5; i++{
		for j := 0; j < 5; j++{
			fmt.Printf("outer line = %v and onner line = %v\n", i, j)
		}
		fmt.Println("------------------------------------")
	}
}

// infinate loop
func excercise_eight() {
	i := 25
	for i > -8{
		fmt.Println(i)
		i--
	}
}


func excercise_nine_a() {
	count := 100
	for {
		if count < 0 {
			break;
		} else if count % 2 == 0 {
			count--
			continue
		}
		
		fmt.Printf("%d\n", count)
		count--;
	}
}
func excercise_nine() {
	i := 25
	for {
		if i == 15{
			break
		}
		fmt.Println(i)
		i--
	}
}

func excercise_ten() {
	xi := [] int {42, 43, 44, 45, 46, 47}

	for i, v := range xi{
		fmt.Printf(" index = %v and value = %v\n", i, v)
	}
}