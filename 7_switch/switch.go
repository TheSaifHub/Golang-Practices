package main

import "fmt"
import "time"

func main() {
	// // switch and case
	// 	i := 2
	// 	switch i {
	// 	case 1:
	// 		fmt.Println("one")
	// 	case 2:
	// 		fmt.Println("two")
	// 	default:
	// 		fmt.Println("others")
	// 	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("It's Weekday")
	default:
		fmt.Println("It's Workdays.")
	}

	// // type assertion and type switch
	// values := func(i interface{}) {
	// 	switch value := i.(type) {
	// 	case int:
	// 		fmt.Println("It is an Integer", value)
	// 	case string:
	// 		fmt.Println("It is a string", value)
	// 	case bool:
	// 		fmt.Println("its a boolean", value)
	// 	default:
	// 		fmt.Println("Other", value)
	// 	}
	// }

	// values(true)

	// // practice with anonymous function and type switching
	// checktype := func(data any) {
	// 	switch t := data.(type) {
	// 	case int:
	// 		fmt.Println("Its an integer:", t)
	// 	case string:
	// 		fmt.Println("Its a string:", t)
	// 	case bool:
	// 		fmt.Println("Its a boolean:", t)
	// 	default:
	// 		fmt.Println("Its an unknown type:", t)
	// 	}
	// }

	// checktype(true)
	// checktype("true")
	// checktype(24)
	// checktype(25.78)

	// |      This pattern is commonly used in:                      |
	// |                                                             |
	// |      ✔ API input handling                                  |
	// |      ✔ JSON decoding                                       |
	// |      ✔ Logging systems                                     |
	// |      ✔ Middleware                                          |
	// |      ✔ Generic utility functions                           |
	// |                                                             |
	// |      Where you don’t know the input type beforehand.        |

	// test := func(data interface{}) {
	// 	switch data.(type) {
	// 	case int:
	// 		fmt.Println("Its an integer type")
	// 	case bool:
	// 		fmt.Println("Its an boolean type")
	// 	case string:
	// 		fmt.Println("Its an string type")
	// 	case float64:
	// 		fmt.Println("Its an float type")
	// 	default:
	// 		fmt.Println("Its an unknown data type")

	// 	}
	// }

	// test("saif")
	// test(24)
	// test(false)
	// test(35.65)

}
