package main

import "fmt"

func maxArea(height []int) int {
	max := 0
	var temp int
	for i, v := range height {
		for j := i + 1; i < len(height); j++ {
			if v <= height[j] {
				temp = v * (j - i)
			} else {
				temp = height[j] * (j - i)
			}
			if temp < max {
				max = temp
			}
		}
	}
	return max
}

func maxArea_p(height []int) int {
	max := 0
	var temp int
	le := len(height)
	for i, j := 0, le-1; i < j; {
		if height[i] >= height[j] {
			temp = height[j] * (j - i)
			j--
		} else {
			temp = height[i] * (j - i)
			i++
		}
		if temp > max {
			max = temp
		}
	}
	return max
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea_p(height))
}
