package main

import (
	"fmt"
	// "slices"
)

func main() {
	// uninitialized slices is nil (means null or empty)
	// var nums []int
	// fmt.Println(nums)
	// fmt.Println(len(nums))

	// making slices
	// var nums = make([]int, 2, 5)
	// cap or capacity -> maximum number of elements can fit, means how many number can fit in slice, 
	// 					  but it doubles when it greater than the initial capacity
	// fmt.Println(cap(nums))
	// fmt.Println(nums)

	// fmt.Println(nums == nil)

	// var nums = make([]int, 2, 5)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// we taken initial capacity 5, but it doubles when it greater than the 5. like above given example

	// nums := []int{}

	// nums = append(nums, 2)
	// nums = append(nums, 1)

	// nums[0] = 3

	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))


	// // copy function (slices)
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 3)
	// var nums2 = make([]int, len(nums))

	// copy(nums2, nums)

	// fmt.Println(nums, nums2)


	// // slice operator
	// nums := []int {1, 2, 3, 4}
	// fmt.Println(nums[1:3])

	// // checking two slices equal or not
	// nums1 := []int {1, 2, 3}
	// nums2 := []int {1, 2}

	// // fmt.Println(nums1 == nums2) // it gives error because slice can only be compared to nil (means empty or null)

	// fmt.Println(slices.Equal(nums1, nums2))

	// // we can also make 2D slices
	nums := [][]int {{1, 3, 4}, {5, 3, 4}}
	// nums = append(nums, []int{2})
	// nums[1] = append(nums[1], 2)
	nums[0] = append(nums[1], 2)
	fmt.Println(nums)
}