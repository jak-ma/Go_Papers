package main

import "fmt"

func main() {
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	//var s string
	//fmt.Println("input:")
	//fmt.Scan(&s)
	//s1 := []rune(s)
	//fmt.Println(RomanToInt(s1, m))
	fmt.Println(RomanToInt([]rune("III"), m))
	fmt.Println(RomanToInt([]rune("MCMXCIV"), m))
	fmt.Println(RomanToInt([]rune("LVIII"), m))
	fmt.Println(RomanToInt([]rune("IX"), m))
	fmt.Println(RomanToInt([]rune("IXIXIX"), m))
}

// lv3
func RomanToInt(s []rune, m map[rune]int) int {
	s = append(s, 'I')
	num := 0
	for i := 0; i < len(s)-1; i++ {
		if m[s[i]] >= m[s[i+1]] {
			num += m[s[i]]
		} else {
			num = num + m[s[i+1]] - m[s[i]]
			i++
		}
	}
	return num
}
