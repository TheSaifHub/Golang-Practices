package main

import "fmt"

// func changeNum(num *int) {
// 	*num = 21
// 	// fmt.Println("Before changing num:", num)
// }

// func main() {
// 	num := 45
// 	changeNum(&num)
// 	fmt.Println("after changed number:", num)
// }



// // pointers practice with swap numbers
// func swap(a *int, b *int) {
// 	temp := *a
// 	*a = *b
// 	*b = temp
// }

// func main() {
// 	x := 5
// 	y := 9
// 	swap(&x, &y)
// 	fmt.Println("x:", x, "y:", y)
// }


// // pointers for changing the string 
// func changeName(name *string) {
// 	*name = "Saif"
// }

// func main() {
// 	myName := "GoLang"
// 	changeName(&myName)
// 	fmt.Println("Name:", myName)
// }



// array element using pointers
func changeFirst(arr *[3]int) {
	arr[0] = 99
}

func main() {
	numbers := [3]int{1, 2, 3}
	changeFirst(&numbers)
	fmt.Println(numbers)
}


