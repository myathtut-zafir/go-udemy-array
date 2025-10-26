package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "www.google.com",
		"aws":    "www.aws.com",
	}
	websites["GitHub"] = "www.github.com"
	delete(websites, "aws")
	fmt.Println(websites)
}