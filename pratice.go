package main

import "fmt"

type product struct {
	id    int
	name  string
	price float64
}

func main2() {
	hobbies := []string{"reading", "swimming", "coding"}
	firstHobby := hobbies[0]
	newHobbies := hobbies[1:3]

	fmt.Println("Hobbies second and third:", newHobbies)
	fmt.Println("Hobbies first:", firstHobby)

	resliece := newHobbies[0:1]
	fmt.Println("Reslice first element:", resliece)

	courseGoals := []string{"Learn Go", "Build Projects"}
	courseGoals[1] = "Master Go"
	fmt.Println("Updated Course Goals:", courseGoals)

	products := []product{
		{id: 1, name: "Laptop", price: 999.99},
		{id: 2, name: "Smartphone", price: 499.99},
	}
	products = append(products, product{id: 3, name: "Tablet", price: 299.99})
	fmt.Println("Product Catalog:", products)

	prices := []float64{19.99, 29.99, 39.99}
	discountedPrices := []float64{17.99, 26.99}
	prices = append(prices, discountedPrices...)
	fmt.Println("All Prices:", prices)

}
