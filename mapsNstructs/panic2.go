package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("deferred")
	panic("bad happened")
	fmt.Println("end")

	//start
	//deferred
	//panic: something bad happened

	// recover returns nil if no panic and vice versa
	fmt.Println("start")
	defer func() {
		if err :=  recover(); err != nil {
			log.Println("Error: ", err)
		}
	}() 
	panic("bad happened")
	fmt.Println("end")

}
