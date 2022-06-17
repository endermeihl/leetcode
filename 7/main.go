package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	num := strconv.Itoa(int(math.Abs(float64(x))))
	l := len(num)
	result := make([]byte, l)
	for i, j := l-1, 0; i >= 0; i, j = i-1, j+1 {
		result[j] = num[i]
	}
	re := string(result)
	in, _ := strconv.Atoi(re)
	if x < 0 {
		if -in < math.MinInt32 {
			return 0
		}
		return -in
	}
	if in > math.MaxInt32 {
		return 0
	}
	return in
}

func reverse_plus(x int) int {
	if x == 0 {
		return x
	}
	var re int
	for x != 0 {
		re = x%10 + re*10
		if re > math.MaxInt32 {
			return 0
		}
		if re < math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	return re
}

func main() {
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse_plus(-123))
}
