package main

import "fmt"

func main() {
	s := []int{1,2,3,4}
	s = append(s, 5)
	fmt.Println(s)

	s1 := make([]int, 5)
	fmt.Println(s1)

	s2 := [6]int {1,2,3}
	fmt.Println(s2)
}