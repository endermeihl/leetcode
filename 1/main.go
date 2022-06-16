package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i, v1 := range nums {
		for j, v2 := range nums[i+1:] {
			if v1+v2 == target {
				return []int{i, i + 1 + j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{-1, -2, -3, -4, -5}
	target := -8
	result := twoSumPlus(nums, target)
	fmt.Println(result)
}

func twoSumPlus(nums []int, target int) []int {
	var hashtable = make(map[int]int)
	for i, v := range nums {
		hashtable[v] = i
	}
	for i, v := range nums {
		j, ok := hashtable[target-v]
		if ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}
