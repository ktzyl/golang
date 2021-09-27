package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
	Phone Phone
}

type Phone struct {
	Country string
	Number string
}

func main() {
	p := Person{
		First: "John",
		Last:  "Doe",
		Age:   30,
		Phone: Phone{
			Country: "1",
			Number: "18818996981",
		},
	}
	q := Person{"Jane","Doe",25,Phone{ "2","15052259830"}}
	fmt.Println(p)
	fmt.Println(q)

	pt := struct {
		X int
		Y int
	}{
		X: 10,
		Y: 20,
	}
	fmt.Println(pt)
}