package main

import "fmt"

// range is used for iterating over data structure
func main() {

	// // for slice
	// nums := []int{6, 7, 8}

	// for i:=0; i<len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// for index, value := range nums {
	// 	fmt.Println(index, value)
	// }

	// // for maps
	// m := map[string]string{"fname":"Saif", "lname":"Ibrahim"}

	// for i, v := range m {
	// 	fmt.Println(i,"=", v)
	// }

	// // for string
	str := "Saif"

	for i, v := range str {
		fmt.Println(i, string(v))
	}
	

}
