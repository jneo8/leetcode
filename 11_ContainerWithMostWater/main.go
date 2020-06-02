package main

func maxArea(height []int) int {
	leftIdx := 0
	rightIdx := len(height) - 1

	maxArea := 0
	for {
		if leftIdx == rightIdx {
			break
		}
		left := height[leftIdx]
		right := height[rightIdx]
		h := 0
		if right < left {
			h = right
			rightIdx--
		} else {
			h = left
			leftIdx++
		}
		area := h * (rightIdx - leftIdx + 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
