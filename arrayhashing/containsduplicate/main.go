package main

import (
	"fmt"
	"sort"
)

func main(){
	
	num := [0]int{}

	slicic := num[:]

	fmt.Print(containsduplicate(&slicic))
}

func containsduplicate(nums *[]int) bool {
	sortedNums := sort.IntSlice(*nums)

	for i := 0; i < len(sortedNums)-1; i ++{
		if sortedNums[i+1] == sortedNums[i]{
			return true
		}
	}
	return false
}