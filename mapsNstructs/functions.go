package main

import(
	"fmt"
)

func main(){
	greeting :="hello"
	name :="stacey"
	sayGreeting(&greeting,&name)
	fmt.Println(name)   //ted



	s :=sum(1,2,3,4,5)
	fmt.println(s)   //15
}


func sum(values ... int) int {
	fmt.Println(values)    // [1 2 3 4 5]
	result :=0

	for _,v :=range values {
		result +=v
	}
return  result
}

func sayGreeting(greeting *string,name *string ){
	fmt.println(*greeting,*name)
	*name="Ted"
	fmt.Println(*name)  //ted 
 
}

