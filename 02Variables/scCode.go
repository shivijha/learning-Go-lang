package main

import "fmt"

const LoginToken string = ";lkjhgvc"

func main() {
	var username string = "Shivani"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggeedIn bool = false
	fmt.Println(isLoggeedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggeedIn)

	// var smallVal unit8 = 255
	// fmt.Println(smallVal)
	// fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.455445112544517677
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var webside = "ShiviCodeLearn.in"
	fmt.Println(webside)

	noOfUser := 300
	fmt.Println(noOfUser)
	fmt.Printf("Variable is of type: %T", noOfUser)

}
