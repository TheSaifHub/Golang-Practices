package main

import "fmt"

// // Closure: Closure remembers the variables which even its outer function and finished execution.

// func counters() func() int {
// 	count := 0

// 	return func() int {
// 		count += 1
// 		return count
// 	}
// }

// func main() {
// 	increment := counters()

// 	for i := range 3 {
// 		// fmt.Println(i)
// 		fmt.Println(i, increment())
// 	}

// 	// fmt.Println(increment())
// }

// // // closure with strings
// func name() func() {
// 	name := "Saif"

// 	return func() {
// 		fmt.Println(name)
// 		// return name
// 	}
// }

// func main() {
// 	printName := name()
// 	printName()

// 	// printName()
// }

// closure with parameters
func multiply(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}

func main() {
	double := multiply(2)
	fmt.Println(double(4))
}


