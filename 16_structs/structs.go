package main

import "fmt"

// type Student struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	s := Student{
// 		Name: "Saif",
// 		Age:  21,
// 	}
// 	fmt.Println(s)
// }

// type Student struct {
// 	Name string
// 	Age  int
// }

// func updateAge(s *Student) {
// 	s.Age = 22
// }

// func main() {
// 	stu := Student{"Rahul", 20}
// 	updateAge(&stu)
// 	fmt.Println("Updated Student:", stu)
// }

// // function returning struct
// type Employee struct {
// 	ID   int
// 	Name string
// }

// func createEmployee() Employee {
// 	return Employee{101, "Amit"}
// }

// func main() {
// 	emp := createEmployee()
// 	fmt.Println(emp)
// }

// slice of structs
type Book struct {
	Title  string
	Author string
}

func main() {
	books := []Book{
		{"Go Basics", "Saif"},
		{"Advanced Go", "Tauseef"},
	}

	for _, b := range books {
		fmt.Println(b)
	}
}
