package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	shivi := User{"Shivi", "shivi@go.dev", true, 16}
	fmt.Println(shivi)
	fmt.Printf("shivi details are: %v", shivi)
	fmt.Printf("name is %v and email is %v", shivi.Name, shivi.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
