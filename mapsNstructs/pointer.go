 package main
 
 import (
	 "fmt"
 )

 type myStruct struct{
	 foo int
 }

 func main(){
	 a :=42
	 b:=&a
     fmt.Println(a,*b)    //42 42

	 a=27
	 fmt.Println(a,*b)      //27 27  
	 
	 *b=14
	 fmt.Println(a,*b)    //14 14



	a: [3]int{1,2,3}
	b:=&a[0]
	c:=&a[1]
	fmt.Printf("%v %p %p\n",a,b,c)  // [1 2 3] address_of_a[0] address_of_a[1] 
  

var ms myStruct
ms= &myStruct{foo:42}
fmt.Println(ms)

}   