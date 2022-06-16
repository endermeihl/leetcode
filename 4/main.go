package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	n := (l1 + l2) % 2
	m := (l1 + l2) / 2
	nums := make([]int, l1+l2)
	i, j, k := 0, 0, 0
	for i < l1 && j < l2 {
		if nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			i++
			k++
		} else if nums1[i] > nums2[j] {
			nums[k] = nums2[j]
			j++
			k++
		} else {
			nums[k] = nums1[i]
			nums[k+1] = nums2[j]
			i++
			j++
			k = k + 2
		}
	}
	for i < l1 {
		nums[k] = nums1[i]
		i++
		k++
	}
	for j < l2 {
		nums[k] = nums2[j]
		j++
		k++
	}

	if n == 1 {
		return float64(nums[m])
	} else {
		return (float64(nums[m]) + float64(nums[m-1])) / 2.0
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Print(findMedianSortedArrays(nums1, nums2))
}
