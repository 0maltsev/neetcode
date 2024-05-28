package main

import (
	"fmt"
)

func main(){
 	nums := []int{4,3,2,1,2}
	fmt.Println(countPrefixArray(nums))
	fmt.Println(countPostfixArray(nums))
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int{
	var resultArr []int
	arraySize := len(nums)
	var result int
	
	prefixArray := countPrefixArray(nums)
	postfixArray := countPostfixArray(nums)

	for i := range nums{
		if i == 0{
			result = postfixArray[arraySize - 2]
		} else  if i == arraySize - 1 {
			result = prefixArray[i - 1]
		} else {
			result = prefixArray[i - 1] * postfixArray[arraySize - 2 -i]
		}
		resultArr = append(resultArr, result)
	}

	return resultArr
}

func countPrefixArray(nums []int) []int{
	prefix := 1

	var prefixArray []int
	for _, element := range nums{
		prefix = prefix * element
		prefixArray = append(prefixArray, prefix)
	}
	return prefixArray
}

func countPostfixArray(nums []int) []int{
	postfix := 1
	arraySize := len(nums)
	var postfixArray []int
	for i := range nums{
		postfix = postfix * nums[arraySize - 1 - i]
		postfixArray = append(postfixArray, postfix)
	}
	return postfixArray
}