package main

import "fmt"

func main() {

	// // if else condition on age
	// age := 20

	// if age <= 18 {
	// 	fmt.Println("You are underage and cannot drive car.")
	// } else {
	// 	fmt.Println("You Can drive car")
	// }

	// // using if, else if and else

	// age := 140
	// if 0 < age && age < 100 {
	// 	if age < 18 {
	// 	fmt.Println("You are Minor.")
	// 	} else if age >= 18 {
	// 		fmt.Println("You are Adult.")
	// 	}
	// } else {
	// 	fmt.Println("Please Enter a valid Age")
	// }

	// // logical operator in for loop
	// var role = "admin"
	// var has_permission = true

	// if role == "admon" || has_permission {
	// 	fmt.Println("You are permitted")
	// } else {
	// 	fmt.Println("You don't have permission")
	// }

	// var role = "admin"
	// var has_permission = false

	// if role == "admin" && has_permission {
	// 	fmt.Println("You are permitted")
	// } else {
	// 	fmt.Println("You don't have permission")
	// }

	// // we can declare variable inside the if construct
	if age := 10; age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 12 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a kid")
	}
}
