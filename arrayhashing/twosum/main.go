package main

import (
	"fmt"
)

func main(){
	nums := []int{3,2,4}
	target := 6

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {return nil}
    
	m := make(map[int]int, len(nums))

	for i := range nums{
		searching := target - nums[i]
		if j,ok := m[searching]; ok{
			return []int{j, i}
		}
		m[nums[i]] = i
	}
	return nil
}
