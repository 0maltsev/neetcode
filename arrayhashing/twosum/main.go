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
	if len(nums) <=1 {return nil}
    m := arrayToHashmap(nums)
	fmt.Println(m)
	for key, value :=range(m){
		first := value[0]
		
		if len(m[key]) > 1{
		m[key] = value[1:]
		} else {
			m[key] = []int{}
		}

		searchingNum := target-key

		if m[searchingNum] != nil && len(m[searchingNum])!=0{
			return []int{first, m[searchingNum][0]}
		}
	}
	return nil
}

func arrayToHashmap(nums []int) map[int][]int {
	m := make(map[int][]int)
	for place, element := range(nums){
		m[element] = append(m[element], place)
	}
	return m
}