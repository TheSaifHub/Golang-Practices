package main

// // function practice on numbers
// func add(a int, b int) int {
// 	return a + b
// }

// func subtract(a, b int) int {
// 	return a - b
// }

// func subtract2() int {
// 	return subtract(10, 5)
// }

// func main() {
// 	addition := add(5, 10)
// 	// subtraction := subtract(10, 5)
// 	subtraction := subtract2()
// 	fmt.Println(addition, subtraction)
// }

// // function practice on strings
// func get_lang() (string, string, string) {
// 	return "Python", "Java", "C++"
// }

// func main() {
// 	lang1, lang2, lang3 := get_lang()
// 	fmt.Println(lang1, lang2, lang3)

// 	// fmt.Println(get_lang())
// }

// // function with multiple return values
// func function(a, b int) (int, int) {
// 	return a/b, a%b
// }

// func main() {
// 	val1, val2 := function(10, 3)
// 	fmt.Println(val1, val2)
// }

// anonymous functions
// func main() {
// 	// // anonymous function declared in a variable
// 	// add := func(a, b int) int {
// 	// 	return a + b
// 	// }

// 	// result := add(5, 5)
// 	// fmt.Println(result)

// 	// anonymous function called immediately
// 	result := func(a, b int) int {
// 		return a + b
// 	}(10, 5)

// 	fmt.Println(result)

// 	// func() {
// 	// 	fmt.Println("Hello Saif!")
// 	// }()
// }

// This is the real world example of anonymous function usage.
// func main() {

// 	// Route: /hello
// 	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "Hello from anonymous function")
// 	})

// 	// Starts server on port 8080
// 	fmt.Println("Server started at http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {

	fn := processIt()
	fn(5)
}
