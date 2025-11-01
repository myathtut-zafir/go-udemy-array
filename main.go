package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	double := transformNumber(&numbers, double)
	fmt.Println(double)
}
func transformNumber(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {    
	return number * 2
}

func tripe(number int) int {
	return number * 3
}
