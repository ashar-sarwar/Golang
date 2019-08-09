//slice is single one means if we refer it somewhere else,
// only one copy is there for every reference

package main 

import (
    "fmt"
)

func main(){ 
 	grade1 := [] int{97,33,54}
    fmt.Printf("Grades: %v , %v, %v",grade1, grade2, grade3)
 

a = make ([]int,3) //(type of obj, length of slice)
b= []int{}
b=append(b,1)
fmt.Println(a)
fmt.Printf("Length : %v\n",len(a))
fmt.Printf("Capacity : %v\n",cap(a))


a=[]int{1,2,3,4,5}
fmt.Println(a)   // [1 2 3 4 5]
b :=append(a[:2],a[3:]...)
fmt.Println(b)      // [1 2 4 5]
fmt.Println(a)    // [1 2 4 5 5]