package main

import "fmt"
// this is function as deep dive

type transformFn func(int) int

func main() {
	numbers := []int{3, 2, 3, 4, 5}
	newFunc := getTransformFn(&numbers)
	double := transformNumber(&numbers, newFunc)

	fmt.Println(double)
}
func transformNumber(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformFn(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}
