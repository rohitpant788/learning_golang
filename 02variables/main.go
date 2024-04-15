package main

import "fmt"

func main() {
	var username string = "rohit"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var smallFloat float64 = 255.123455555555555555
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

}
https://golang.org/ref/spec#Numeric_types
