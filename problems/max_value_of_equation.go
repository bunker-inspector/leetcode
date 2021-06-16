package problems

import (
	"math"
)

func MaxValueOfEquation(points [][2]int, k int) int {
	if len(points) == 0 || k == 0 {
		return 0
	}

	max := math.MinInt64
	stack := [][2]int{}

	for _, point := range points {
		for len(stack) > 0 && point[0]-stack[0][0] > k {
			stack = stack[1:]
		}
		if len(stack) > 0 {
			val := point[0] + point[1] + stack[0][1]
			if val > max {
				max = val
			}
		}
		for len(stack) > 0 && point[1]-point[0] > stack[len(stack)-1][1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, [2]int{point[0], point[1] - point[0]})
	}
	return max
}
