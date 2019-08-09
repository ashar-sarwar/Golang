package main

import (
	"fmt"
)

func main(){

	switch 5{
	case 1,2,5:
		fmt.Println("1,2 or 5")
		fallthrough   // moves to below case without any logic
	case 2,4,6:
		fmt.Println("2, 4 or 6")
	default:
		fmt.Println("any num")
	}



var i interface{}=1
switch i.(type){
case int:
	fmt.Println("Int")
case float64:
	fmt.Println("float")
}
default:
	fmt.Println("any type")

	}  