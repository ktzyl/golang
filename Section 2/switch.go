package main

import (
	"fmt"
)

func main() {
	a := 20

	switch a {
	case 10:
		fmt.Println("10")
	case 9:
		fmt.Println("9")
	case 1:
		fmt.Println("1")
	default:
		fmt.Println("None of these")
	}
}
