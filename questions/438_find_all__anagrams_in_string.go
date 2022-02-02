package leetcode

import "fmt"

func findMap(s string) map[string]int {
	resMap := make(map[string]int)

	for i := 0; i < len(s); i++ {
		resMap[string(s[i])]++
	}

	return resMap
}

func isAnagram(s map[string]int, p map[string]int) bool {
	for key, element := range p {
		if element != s[key] {
			return false
		}
	}

	return true
}

func findAnagrams(s string, p string) []int {
	lenS := len(s)
	lenP := len(p)

	if lenS < lenP {
		return []int{}
	}

	resMap := findMap(p)
	tempMap := findMap(s[:len(p)])
	var res []int

	if isAnagram(resMap, tempMap) {
		res = append(res, 0)
	}

	for i := 1; i <= lenS-lenP; i++ {
		tempMap[string(s[i-1])]--
		tempMap[string(s[i+lenP-1])]++

		if isAnagram(resMap, tempMap) {
			res = append(res, i)
		}
	}
	return res
}

func Q_438() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}
