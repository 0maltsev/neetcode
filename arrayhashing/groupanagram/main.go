package main

import (
	"fmt"
)

type Key [26]byte

func main(){
	strs := []string{"eat","tea","tan","ate","nat","bat"}
	fmt.Println(convertGrouppedAnagramsToArray(strs))
}

func stringToKey(str string) Key{
	var k Key
	for _, letter := range str{
		k[letter - 'a']++
	}
	return k
}

func groupAnagrams(strs []string) map[Key][]string{
	m := make(map[Key][]string)
	for _, str := range strs{
		key := stringToKey(str)
		m[key] = append(m[key], str)
	}
	return m
}

func convertGrouppedAnagramsToArray(strs []string) [][]string{
	var result [][]string
	m := groupAnagrams(strs)
	for _, value := range m{
		result = append(result, value)
	}
	return result
}