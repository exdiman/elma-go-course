package main

import "fmt"

func square() func() int {
	x := 0
	return func() int {
		x++
		return x * x
	}
}

func main() {
	count := 10
	iter := square()
	for i := 1; i <= count; i++ {
		fmt.Println(iter())
	}
}
