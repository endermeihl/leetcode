package main

import (
	"container/list"
	"fmt"
)

func isValid(s string) bool {
	var dp map[string]string = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	le := len(s)
	if le%2 != 0 {
		return false
	}
	temp := list.New()
	for i := 0; i < le; i++ {
		if string(s[i]) == "(" || string(s[i]) == "{" || string(s[i]) == "[" {
			temp.PushBack(string(s[i]))
		} else {
			if temp.Len() <= 0 {
				return false
			}
			if temp.Back().Value != dp[string(s[i])] {
				return false
			} else {
				temp.Remove(temp.Back())
			}
		}
	}
	return temp.Len() <= 0
}

func main() {
	s := "}{"
	fmt.Println(isValid(s))
}
