package main

import (
	"fmt"
)

func main() {

	s := sum(1, 2, 3, 4, 5)
	fmt.println(*s) //15
}

func sum(values ...int) *int  {
	fmt.Println(values) // [1 2 3 4 5]
	result := 0

	for _, v := range values {
		result += v
	}   
	return &result
	
}
