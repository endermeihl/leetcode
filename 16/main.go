package main

import (
	"fmt"
	"math"
)

func threeSumClosest_e(nums []int, target int) int {
	le := len(nums)
	max := math.MaxInt32
	var result int
	for i := 0; i < le-2; i++ {
		for j := i + 1; j < le-1; j++ {
			for k := j + 1; k < le; k++ {
				temp := nums[i] + nums[j] + nums[k]
				if int(math.Abs(float64(temp-target))) < max {
					result = temp

				}
			}
		}
	}
	return result
}

func threeSumClosest(nums []int, target int) int {
	sort(&nums)
	le := len(nums)
	max := math.MaxInt32
	var result int
	for i := 0; i < le; i++ {
		start := i + 1
		end := le - 1
		for start < end {
			temp := nums[start] + nums[end]
			if abs(temp+nums[i]-target) < max {
				result = temp + nums[i]
				max = abs(temp + nums[i] - target)
			}
			if temp+nums[i] > target {
				end--
			} else if temp+nums[i] < target {
				start++
			} else {
				return result
			}
		}
	}
	return result
}

func abs(num int) int {
	return int(math.Abs(float64(num)))
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	sort(&nums)
	fmt.Println(nums)
	fmt.Println(threeSumClosest(nums, target))
}

func quickSort(nums *[]int, left int, right int) {

}

func sort(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		for j := i + 1; j < len(*nums); j++ {
			if (*nums)[i] > (*nums)[j] {
				temp := (*nums)[i]
				(*nums)[i] = (*nums)[j]
				(*nums)[j] = temp
			}
		}
	}
}
