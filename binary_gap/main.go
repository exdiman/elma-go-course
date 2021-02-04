package main

import "fmt"

func Solution(N int) (max int) {
	var count int

	if N < 0 {
		N = -N
	}
	for N > 0 {
		if N&1 == 0 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
		N = N >> 1
	}

	return
}

func main() {
	var number int
	fmt.Println("Введите целое число:")
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%b\n", number)
	fmt.Println(Solution(number))
}
