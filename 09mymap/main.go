package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcom to map world")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Pyhton"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all lanuages: ", languages)

	//LOOPS ARE INTERESTING IN GO LANG
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

	//FOR ONLY VALUE
	for _, value := range languages {
		fmt.Printf("for key v, value is %v\n", value)
	}
}
