package main

import "fmt"

var results []string

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	le := len(digits)
	if le == 0 {
		return []string{}
	}
	backtrack(digits, 0, "")
	return results

}

func backtrack(digits string, index int, result string) {
	if index == len(digits) {
		fmt.Println(result)
		results = append(results, result)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		le := len(letters)
		for i := 0; i < le; i++ {
			backtrack(digits, index+1, result+string(letters[i]))
		}

	}
}

func main() {
	digits := "22"
	fmt.Println(letterCombinations(digits))
}
