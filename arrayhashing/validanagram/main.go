package main

import(
	"fmt"
	"reflect"
)

func main(){
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	ms := convertStringToHashMap(s)
	mt := convertStringToHashMap(t)
	return reflect.DeepEqual(ms, mt)
}

func convertStringToHashMap(s string) map[rune]int{
	m := make(map[rune]int)
	for _, letter := range(s){
		m[letter] = m[letter]+1
	}
	return m
}