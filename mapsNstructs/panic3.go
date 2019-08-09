package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")

	defer func() {
	
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err) // re-panic for higher functions in stack to stop execution  
		}
	}()
	panic("bad happened")
	fmt.Println("panikckckkc")

}
 