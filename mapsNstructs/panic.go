package main

import(
	"fmt"
)

//panic happens after the defer funct is executed
//panic is just like exception in JS     
  
         func main(){
		a,b :=1,0
		ans := a/b
		fmt.Println(ans)
		  //panic : runtime error: integer divided by zero  

		  fmt.Println("satrt")
		  panic("buhund")
		  fmt.Println("end")

		  //start
          // panic: something bad happened
  }

 