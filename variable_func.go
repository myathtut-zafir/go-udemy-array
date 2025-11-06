package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(numbers)
	fmt.Println("Sum:", sum)

}

func sumup(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
func sumupManyArg(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
