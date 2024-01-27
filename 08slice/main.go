package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to slice world")

	//var fruitList = []string{"Apple", "Tomato", "Peach"}
	// fmt.Println("Types of fruitlist is %T\n", fruitList)

	// fruitList = append(fruitList, "Mango", "Banana")
	// fmt.Println(fruitList)
	// fruitList = append(fruitList[1:4])
	// fmt.Println(fruitList)

	// highScore := make([]int, 4)
	// highScore[0] = 234
	// highScore[1] = 259
	// highScore[2] = 232
	// highScore[3] = 334

	// highScore = append(highScore, 555)
	// fmt.Println(highScore)
	// sort.Ints(highScore)
	// fmt.Println(highScore)

	// fmt.Println(sort.IntsAreSorted(highScore))

	//HOW TO REMOVE ELEMENT IN SLICE
	course := []string{"reactjs", "javascipt", "swift", "python", "ruby"}
	fmt.Println(course)
	index := 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}
