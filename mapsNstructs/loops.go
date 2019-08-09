  package main
  
  import (
	  "fmt"
  )

  func main(){
	  statePop :=map[string]int{
		  "california":34243,
		  "ohio":324245,
		  "new jersey":54322,
	  }

	  for k,_ range statePop{
		  fmt.Println(k)
	  }
  }