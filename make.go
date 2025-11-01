package main

import "fmt"

type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)

}

func mainH() {
	userName := make([]string, 2, 5)
	userName = append(userName, "Alice", "Bob", "Charlie")

	fmt.Println(userName)
	// courseRatings := make(map[string]float64, 3)
	courseRatings := make(floatMap, 3)
	courseRatings["Math"] = 4.5
	courseRatings["Science"] = 4.7
	courseRatings["History"] = 4.3
	courseRatings.output()

	for index, value := range userName {
		fmt.Println("Index:", index, "Value:", value)
	}
	// for course, rating := range courseRatings {
	// 	fmt.Printf("Course: %s, Rating: %.2f\n", course, rating)
	// }

}
