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
 	for _, element := range(*nums){
 		m[element]++
 		if m[element] > 1{
 			return true
 		}
 	}
 	return false
}