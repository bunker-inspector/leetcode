package problems

func TrappedRainwater(heights []int) int {
	if len(heights) < 3 {
		return 0
	}
	leftmax, rightmax := heights[0], heights[len(heights)-1]
	left, right := 1, len(heights)-2

	water := 0
	for left <= right {
		leftmax = max(heights[left], leftmax)
		rightmax = max(heights[right], rightmax)

		if leftmax < rightmax {
			water += leftmax - heights[left]
			left++
		} else {
			water += rightmax - heights[right]
			right--
		}
	}
	return water
}
