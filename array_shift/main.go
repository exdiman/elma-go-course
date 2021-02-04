package main

import "fmt"

/*
func cyclicRightShift(arr []int, times uint) []int {
	length := len(arr)
	if length == 0 {
		return arr
	}

	shift := length - int(times)%length
	if shift == 0 || shift == length {
		return arr
	}

	return append(arr[shift:], arr[:shift]...)
}
*/

/*
func cyclicRightShift2(arr []int, times uint) []int {
	length := len(arr)
	if length == 0 {
		return arr
	}

	shift := length - int(times)%length
	if shift == 0 || shift == length {
		return arr
	}

	return append(append(make([]int, 0, length), arr[shift:]...), arr[:shift]...)
}
*/

func cyclicRightShiftByPointer(arr *[]int, times uint) {
	length := len(*arr)
	if length == 0 {
		return
	}

	shift := length - int(times)%length
	if shift == 0 || shift == length {
		return
	}

	*arr = append((*arr)[shift:], (*arr)[:shift]...)
}

/*
func cyclicRightShiftByPointer2(arr *[]int, times uint) {
	length := len(*arr)
	if length == 0 {
		return
	}

	shift := length - int(times)%length
	if shift == 0 || shift == length {
		return
	}
	buf := make([]int, shift, shift)
	copy(buf, (*arr)[:shift])
	copy(*arr, (*arr)[shift:])
	copy((*arr)[length-shift:], buf)
}
*/

func main() {
	var times uint = 6
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println("times =", times)
	cyclicRightShiftByPointer(&arr, times)
	fmt.Println(arr)
}
