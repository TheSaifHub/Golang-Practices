package main

import "fmt"

func main() {
	// for -> it is only construct in golang for looping
	

	// // while loop (by the help of using for loop)
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i+=1
	// }

	// // infinite loop
	// for {
	// 	println("1")
	// }

	// // classic for loop

	// for i:=1; i<=5; i++ {
	// 	fmt.Println(i)
	// }

	// // using of continue and break
	// for i:=1; i<=5; i++ {
	// 	if i==2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for i:=1; i<=5; i++ {
	// 	if i==2 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// // use of range
	for i:=range 4 {
		fmt.Println(i)
	}
}