package main

import "fmt"

func romanToInt(s string) int {
	var roman = make(map[string]int)
	roman["I"] = 1
	roman["V"] = 5
	roman["X"] = 10
	roman["L"] = 50
	roman["C"] = 100
	roman["D"] = 500
	roman["M"] = 1000
	le := len(s)
	var result int
	pre := roman[string(s[0])]
	for i := 1; i < le; i++ {
		num := roman[string(s[i])]
		if num >= pre {
			result = result - pre
		} else {
			result = result + pre
		}
		pre = num
	}
	result = result + pre
	return result
}

func main() {
	s := "CIX"
	fmt.Println(romanToInt(s))
}
