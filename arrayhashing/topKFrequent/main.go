package main

import(
	"fmt"
)

func main(){
	nums := []int{1, 1, 1, 2, 2, 3, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {

	var result []int

    m := convertArrayToHashMap(nums, k)

	fmt.Print(m)
	
	for i := len(nums); i > 0; i--{
		valArr := m[i]
		
		for _, element := range valArr{
			if k > 0 {
				result = append(result, element)
				k--
			} else {
				return result
			}
		}
	} 
	return result
	
}

func convertArrayToHashMap(nums []int, k int) map[int][]int{
	m := make(map[int]int)
	n := make(map[int][]int)

	for i:=0; i <= k; i++{
		n[i] = []int{}
	}

	for _, num := range nums{
		m[num] ++
	}

	for key, value := range m {
		n[value] = append(n[value], key)
	}
	return n
}
