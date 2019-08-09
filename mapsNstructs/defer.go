package main

import(
	"fmt"
)

//defer is funct which is executed after the function is executed and BEFORE IT IS RETURNED

func main(){
	fmt.Println("start")
	defer fmt.Println("middle")
	fmr.Println("end")
}

// start
//end
//middle

a:="start"
defer fmt.Println(a) // this will work after func is executed but the value in parameter would be "start" not "end",
a:="end"

/////////////////////////////////////////////////////////////////////////////////////////////////////////


import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	res,err := http.Get("http://anysite.com")

	if err!=nil{
		log.Fatal(err)
	}
	defer res.Body.Close()

	robots,err := ioutil.ReadAll(res.Body)
	if err !=nil{
		log.Fatal(err)
	}

	fmt.Printf("%s",robots)
} 
 
