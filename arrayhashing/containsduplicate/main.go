package main

import (
	"fmt"	
)

func main(){
	
	num := [0]int{}

	slicic := num[:]

	fmt.Print(containsduplicate(&slicic))
}

func containsduplicate(nums *[]int) bool {
	m := make(map[int]int)
	for i:= 0; i < len(*nums); i++ {
		m[(*nums)[i]]++
		if m[(*nums)[i]] > 1{
			return true
		}
	}
	return false
}