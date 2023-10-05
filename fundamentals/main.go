package main
import (
	"fmt"
   // "fundamentals/furtherexplored"
)

type student struct{
	first string 
	id int
}

var x int = 5

func setnames (name string, id int) (s student){
	var data student 
	data.first = name
	data.id = id
	return data
}

func main() {
	var s  student = setnames( "sree", 10)
	fmt.Printf(" name = %s and id = % d", s.first, s.id)
	cust_print()
	s.sayhello()

	if k := 5; k > 2{
		fmt.Println(" \nk is greater than 5")
	}

	m := 7

	if m > 5 {
		fmt.Println("m is greater than 5")
	}

	fmt.Println("for loop demo")

	for i := 1; i < 5; i++ {
		fmt.Printf("\n%d", i)
	}

	furtherexplored.Facscinating()
}

func cust_print(){
	fmt.Printf(" x = %d", x)
}


func (s student) sayhello(){
	fmt.Printf("hi my name is ", s.first)
}