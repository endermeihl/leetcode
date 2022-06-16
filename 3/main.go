package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	max := 1
	len := len(s)
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			k, result := check([]byte(s[i:j]), s[j])
			if result {
				i = i + k + 1 //子字符串返回的需要等效为原字符串位置，并且将左坐标移到该位置
			} else {
				maxtemp := j - i + 1
				if max < maxtemp {
					max = maxtemp
				}
			}
		}
	}
	return max
}

//判断字符是否在子字符串内
//在就返回true和他的位数
//不在返回false
func check(s []byte, t byte) (int, bool) {
	for k, v := range s {
		if v == t {
			return k, true
		}
	}
	return 0, false
}

func main() {
	fmt.Print(lengthOfLongestSubstring("aabcddd"))
}
