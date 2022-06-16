package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	l := len(s)
	if l == numRows || numRows <= 1 {
		return s
	}
	dp := make([][]byte, numRows)
	result := make([]byte, l)
	change := true
	for i := range dp {
		dp[i] = make([]byte, l)
	}
	q, t := 0, 0
	for _, v := range []byte(s) {
		if change {
			dp[t][q] = v
			t++
		} else {
			q++
			t--
			dp[t][q] = v
		}
		if t%numRows == 0 && change {
			change = false
			t--
		} else if t == 0 && !change {
			change = true
			t++
		} else if (t-1)%numRows == 0 && !change {
			change = true
			q++
			t--
		}
	}
	w := 0
	for _, v1 := range dp {
		for _, v2 := range v1 {
			if v2 != 0 {
				result[w] = v2
				w++
			}

		}
	}
	return string(result)
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3

	fmt.Println(convert(s, numRows))
	fmt.Println(convert_p2(s, numRows))
}

func convert_p2(s string, numRows int) string {
	l := len(s)
	if l == numRows || numRows <= 1 {
		return s
	}
	change := true
	t := 0
	dp := make([][]byte, numRows)
	for _, v := range []byte(s) {
		if change {
			dp[t] = append(dp[t], v)
			t++
		} else {
			t--
			dp[t] = append(dp[t], v)
		}
		if t%numRows == 0 && change {
			change = false
			t--
		} else if t == 0 && !change {
			change = true
			t++
		} else if (t-1)%numRows == 0 && !change {
			change = true
			t--
		}
	}
	return string(bytes.Join(dp, nil))
}

//可以提升空间
//1.精简内部逻辑
//2.优化存储空间
//3.直接返回结果字段
