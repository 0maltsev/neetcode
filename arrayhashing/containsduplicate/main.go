package main

import (
	"fmt"
	"sort"
)

func main(){
	
	num := [4]int{1,2,3,1}

	slicic := num[:]

	fmt.Print(containsduplicate(slicic))
}

func containsduplicate(nums []int) bool {
	sort.Ints(nums)
	fmt.Println(nums)

	for i := 0; i < len(nums)-1; i ++{
		if nums[i+1] == nums[i]{
			return true
		}
	}
	return false
}