package main

import "fmt"

func main() {
	s := "My String"
	sptr := &s
	fmt.Println(s)
	fmt.Println(*sptr)

	sptr1 := new(string)
	fmt.Println(*sptr1)

    s2 := "My string2"
	var sptr2 *string
	sptr2 = &s2
	fmt.Println(*sptr2)
}