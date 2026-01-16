package main

import (
	"fmt"
	"maps"
)

func main() {

	// // creating map
	// m := make(map[string]string)

	// // setting element
	// m["name"] = "Saif"
	// m["language"] = "golang"

	// // accessing element
	// fmt.Println(m["name"], m["language"])

	// // if we accessing the key that doesnot exists then it returns empty value
	// fmt.Println(m["gender"])  // it gives the null or empty string

	// // deleting elements
	// delete(m, "language")

	// fmt.Println(m)


	// // creating map with another method
	// m := map[string]int{"age":21, "price":240}
	// fmt.Println(m)

	// // cheking that key value exist or not.
	// value, existence := m["ae"]

	// fmt.Println(value)

	// if existence {
	// 	fmt.Println("Its exist.")
	// } else {
	// 	fmt.Println("Its not exist.")
	// }

	// checking two maps equal or not
	m1 := map[string]int{"age":21, "price":240}
	m2 := map[string]int{"age":21, "price":240}

	fmt.Println(maps.Equal(m1, m2))

}