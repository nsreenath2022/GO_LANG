package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("Hello world\n")

	var x = 15
	var y = 13
	if x < 16 && y == 13 {
		fmt.Println("Hehehe")
	} else  if  x == 15 || y == 7{
		fmt.Println("?????")
	} else {
		fmt.Println("hey man")
	}

	x = 40
	if z := 2 * rand.Intn(x); x < z {
		fmt.Println("what you expect")
	} else {
		fmt.Println(" may not be true")
	}

	switch x{
	case 40:
		fmt.Println("x == 40")
		
    case 42:
		fmt.Println("x == 42")
		
	default:
		fmt.Println("default case for x")
	}
/*
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(){
		ch1 <- 41
	}()

	go func(){
		ch2 <- 42
	}()


	select {
	case v1:= <-ch1:
		fmt.Println("channel one: ", v1)
    case v2:= <-ch2:
		fmt.Println("chanllel two: ", v2)
	}*/

	var m = 10

	for i := 0; i < 10; i++ {
		fmt.Println(m)
	}

	var i = 0
	for i < m {
		fmt.Printf("hehehehe")
		i++
	}

	i = 0
	for {
		if i > 10 {
			break
		}
		i++
		fmt.Printf("umma\n")
	}



	fmt.Printf("even numbers")

	for i = 0 ; i <= m; i++{
		if i % 2 == 0{
			fmt.Println(i)
		}
	}


	// nested loops

	for i =0; i < m; i++{
		for j := i; j <m; j++{
			fmt.Printf("%d <-> %d", j, i)
		}
		fmt.Printf("\n")
	}

	var xi = [] int {34,12,45,39,80}

	for i, v := range xi {
		fmt.Printf("index = %v and value = %v\n", i, v)
	}

	names  := [] string {"chaitra", "vaishaka", "jesta", "aashada", "sgrerava"}

	for i, v := range names{
		fmt.Printf("index = %v and name = %v\n", i, v)
	}
	fmt.Printf("end\n")

	prices := []float32{12.98, 89.12, 0.06, 78.97}
	for i, v := range prices {
		fmt.Printf(" index = %d and price = %f\n", i, v)
	}


	//maps
	 mp := map[string]int{
		"Sree": 6,
		"Swa": 9,
		"Sow": 8,
		"Amma": 5,
		"Apps": 7,
	}

	// print vals
	for key, val:= range mp {
		fmt.Printf(" key = %s and value = %d\n", key, val)
	}


	q := 89 / 23
	r := 69 % 3

	fmt.Printf(" q = %d and r = %d\n", q, r)


	year := [4]int {1984, 1988, 1923, 1900}


	for val, year := range year {
		fmt.Printf(" year = %d  idx = %d \n", year, val)
	}


	xs := [] string {"hello", "world"}
	fmt.Println(xs)
	
	for _, val := range xs{
		fmt.Println(val)
	}
	
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(len(xs))

	fmt.Println("Demo of for loop using index\n")

	for x:=0; x <len(xs); x++{
		fmt.Println(xs[x])
	}


	// appending to a slice
	var slice = [] int {34,12,45,39,80}
	fmt.Println(slice)
	fmt.Println("append starts\n")
	slice = append(slice, 96)
	fmt.Println(slice)
	slice = append(slice, 45)
	fmt.Println(slice)

	slice = append(slice ,98.00)
	fmt.Println(slice)

	slice = append(slice , 34,78,102,431,456,879)
	fmt.Println(slice)
	fmt.Printf("slaice = %#v\n", slice)
	// slicing a slice
	// [inclusive:exclusive]
	fmt.Printf("slice = %#v\n", slice[0:3])

	fmt.Printf("--------------------\n")
	fmt.Printf("original slice = %#v\n", slice)
	fmt.Printf("slice = %#v\n", slice[1:5])

	fmt.Printf("--------------------\n")
	fmt.Printf("original slice = %#v\n", slice)
	fmt.Printf("slice = %#v\n", slice[6:])

	fmt.Printf("--------------------\n")
	fmt.Printf("original slice = %#v\n", slice)
	fmt.Printf("slice = %#v\n", slice[:4])

	fmt.Printf("--------------------\n")
	fmt.Printf("original slice = %#v\n", slice)
	fmt.Printf("slice = %#v\n", slice[5:5])


	// delting from slice
	fmt.Printf("--------------------\n")
	fmt.Printf("original slice = %#v\n", slice)
	slice = append(slice[:3], slice[4:]...)
	fmt.Printf("slice = %#v\n", slice)

	// length and capacity
	// delting from slice
	fmt.Printf("--------------------\n")
	fmt.Printf("original slice = %#v\n", slice)
	fmt.Printf("slice  len =%v and capacity = %v\n", len(slice), cap(slice))

	//plice := make([]int , 0, 10)

	plice := slice
	fmt.Println(plice)


	// structs

	type Person struct{
		first string
		last string
		age int
	}


	type Pets struct{
		Person
		types int
		variety int
	}
	p1 := Person{
		first : "Sreenath",
		last: "XYZ",
		age: 32,
	} 

	pets := Pets{
		Person: Person{
			first : "Sreenath",
			last: "XYZ",
			age: 32,
		},
		types: 5,
		variety: 6,
	}

	fmt.Println(p1)
	fmt.Println(pets)

	p2:= p1
	fmt.Println(p2)

	p2.age = 99
	fmt.Println(p2)

	p2.first = "kantara"
	fmt.Println(p2)



	fmt.Printf("sum = %d", doOp(1, 2, add))
	fmt.Printf("sum = %d", doOp(1, 2, sub))
}


func add (a int, b int)int{
	return  a + b
}

func sub (a int, b int)int{
	return  a - b
}


func addT[T numbers](a,b T) T{
	return a + b
}

func doOp(a int, b int, f func(int,int)int) int{
	return f(a, b)
}