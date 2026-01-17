package main

import "fmt"

func counters() func() int {
	count := 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counters()

	for i := range 3 {
		// fmt.Println(i)
		fmt.Println(i, increment())
	}
	
	// fmt.Println(increment())
}