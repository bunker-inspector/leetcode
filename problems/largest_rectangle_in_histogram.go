package problems

func LargestRectangle(heights []int) int {
	var leftbounds []int
	var rightbounds []int
	var tracker []int

	for i := 0; i < len(heights); i++ {
		for {
			if len(tracker) == 0 {
				leftbounds = append(leftbounds, -1)
				tracker = append(tracker, i)
				break

			}
			peek := tracker[len(tracker)-1]
			if heights[peek] < heights[i] {
				leftbounds = append(leftbounds, peek)
				tracker = append(tracker, i)
				break
			}
			tracker = tracker[0 : len(tracker)-1]
		}
	}
	tracker = []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		for {
			if len(tracker) == 0 {
				rightbounds = append(rightbounds, len(heights))
				tracker = append(tracker, i)
				break

			}
			peek := tracker[len(tracker)-1]
			if heights[peek] < heights[i] {
				rightbounds = append(rightbounds, peek)
				tracker = append(tracker, i)
				break
			}
			tracker = tracker[0 : len(tracker)-1]
		}
	}

	max := 0
	for i := 0; i < len(heights); i++ {
		curr := ((rightbounds[len(heights)-i-1] - leftbounds[i]) - 1) * heights[i]
		if curr > max {
			max = curr
		}
	}
	return max
}
