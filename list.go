package main

import "fmt"

func main() {
	prices := []float64{19.99, 29.99, 39.99}
	var productNames []string
	productNames = append(productNames, "hi", "hello", "yes")
	// fmt.Println("Hello, World!", prices[0])
	// fmt.Println("Hello, World!", productNames)
	featurePrices := prices[1:3]
	featurePrices = prices[:3]
	fmt.Println("Hello, World!", featurePrices)
	fmt.Println("Hello, World!", len(featurePrices))
}
