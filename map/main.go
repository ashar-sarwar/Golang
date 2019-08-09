package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":  "abcd",
	// 	"blue": "uady9",
	// }

	colors := make(map[int]string)

	colors[10] = "helloi"
	colors[40] = "43r213"
	colors[32] = "2423"
	//delete(colors, 10)
	//fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[int]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
