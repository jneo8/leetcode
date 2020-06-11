package main

import (
	"sort"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[len(nums)-1]

	for i := 0; i < len(nums)-2; i++ {
		a, left, right := nums[i], i+1, len(nums)-1
		for left < right {
			sum := a + nums[left] + nums[right]
			if target == sum {
				return sum
			}
			if abs(target-sum) < abs(target-closest) {
				closest = sum
			}
			if target-sum > 0 {
				left++
			} else {
				right--
			}
		}
	}
	return closest
}
