package leetcode

import "fmt"

func numParser(s string) int {
	switch s := s[0]; s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		fmt.Println("Error while parsing", s)
		return -1

	}
}

func isNegeted(s1 string, s2 string) bool {
	return numParser(s1) < numParser(s2)
}

func romanToInt(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		currentStrings := string(s[i])

		if i != len(s)-1 && isNegeted(currentStrings, string(s[i+1])) {
			result += numParser(string(s[i+1])) - numParser(currentStrings)
			i++
			continue
		}

		result += numParser(currentStrings)
	}
	return result
}

func Q_13() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
