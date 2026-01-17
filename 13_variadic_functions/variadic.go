package main

import "fmt"

// func add(nums ...int) int {
// 	total := 0

// 	for _, num := range nums {
// 		total = total + num
// 	}

// 	return total
// }

// func main() {
// 	number := []int{3, 4, 5, 6, 10}
// 	result := add(number...)
// 	// result := add(3, 4, 5, 5, 5, 5, 10, 10, 10, 10)
// 	fmt.Println(result)
// }








// func fn(nums ...string) string {
// 	total := ""

// 	for _, num := range nums {
// 		total = total + num
// 	}

// 	return total
// }

// func main() {
// 	result := fn("saif", "sufiyan", "tauseef")
// 	fmt.Println(result)
// }








// // variadic function with different different type params 
// func data(values ...interface{}) {  // at the place of interface{} we can also use any
// 	for _, val := range values {
// 		fmt.Println(val)
// 	}
// }

// func main() {
// 	data("saif", 21, 9170055886, 40.6, true)
// }




// // variadic and anonymous function together
// func main() {
// 	function := func(values ...interface{}) {
// 		for _, val := range values {
// 			fmt.Println(val)
// 		}
// 	}

// 	function("saif", 24)
// }




// variadic function with slices
func data(values ...interface{}) {
	fmt.Println(values...)
}

func main() {
	slice := []any{"saif", 35, 78.098, false}
	data(slice...)
}