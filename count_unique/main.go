package main

import "fmt"

func countUnique(arr []int) int {
	m := make(map[int]bool)

	for _, number := range arr {
		m[number] = true
	}

	return len(m)
}

func main() {
	arr := []int{1, 1, 1, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(countUnique(arr))
}
