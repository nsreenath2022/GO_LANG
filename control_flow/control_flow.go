package main

import (
	"fmt"
)

func main () {
	print_even_numbers()
	matrix := [][]int{{1,2,3,4,5,6}, {3,4,5,6,7,8}, {5,6,1,9,4,7}}
	print_matrix(matrix)
}

func print_even_numbers(){
	for i := 0; i < 20; i++{
		if i % 2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

func print_matrix (matrix [][] int){

	//row1:= matrix[0]
	//row2:= matrix[1]

	columns  := len(matrix[0])
	rows := len(matrix)

	fmt.Printf("\nnumber of rows = %v and number of columns = %v\n", rows, columns)
	
	for i := 0; i  < rows; i++{
		for j := 0; j < columns; j++{
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println("\n")
	}

}