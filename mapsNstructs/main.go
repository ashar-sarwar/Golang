//map cannot guarantee the order of items
//map only has one copy like slices
package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{

		"california": 28379283,
		"iduiiwd":    39829173,
	}

	m := map[[3]int]string{}
	fmt.Println(statePopulations, m)

pop,ok :=statePopulations["california"]
fmt.Println(pop, ok)  // 28379283 true

}


 