package main

import "fmt"

func main() {
	m := map[string]string{}
	m["first"] = "John"
	m["last"] = "doe"
	fmt.Println(m)

	m1 := make(map[string]string)
	m1["first"] = "Jane"
	m1["last"] = "Ai"
	fmt.Println(m1["first"])
}