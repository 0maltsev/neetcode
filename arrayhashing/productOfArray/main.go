package main

import (
	"fmt"
)

func main(){
 	nums := []int{4,3,2,1,2}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int{
	result := make([]int, len(nums))

	prefix := 1
	for i, element := range nums{
		result[i] = prefix
		prefix *= element
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}
	return result
}

