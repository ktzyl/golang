package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int
	for _, a := range os.Args[1:] {
		i, err := strconv.Atoi(a)
		if err != nil {
			panic(fmt.Sprintf("Invalid Value: %v", err))
		}
		sum += i
	}
	fmt.Printf("Sum = %v\n", sum)
}

strconv