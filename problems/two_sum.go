package problems

func TwoSum(nums []int, target int) []int {
	seenVals := make(map[int]int)
	for idx, val := range nums {
		needed := target - val
		seenVal, present := seenVals[needed]
		if present {
			return []int{seenVal, idx}
		}
		seenVals[val] = idx
	}
	return []int{0, 0}
}
