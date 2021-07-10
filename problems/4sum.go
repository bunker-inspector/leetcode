package problems

import (
	"sort"
)

type pair struct {
	idx1 int
	idx2 int
}

func FourSum(nums []int, target int) [][4]int {
	pairs := make(map[int][]pair)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			pairsum := nums[i] + nums[j]
			if curr, ok := pairs[pairsum]; ok {
				pairs[pairsum] = append(curr, pair{i, j})
			} else {
				pairs[pairsum] = []pair{{i, j}}
			}
		}
	}

	results := make(map[[4]int]bool)
	for sum, pairList := range pairs {
		complementNum := target - sum
		for _, pair := range pairList {
			if complementPairs, ok := pairs[complementNum]; ok {
				for _, complementPair := range complementPairs {
					if pair.idx1 != complementPair.idx1 &&
						pair.idx1 != complementPair.idx2 &&
						pair.idx2 != complementPair.idx1 &&
						pair.idx2 != complementPair.idx2 {
						newResult := [4]int{
							nums[pair.idx1],
							nums[pair.idx2],
							nums[complementPair.idx1],
							nums[complementPair.idx2],
						}
						sort.Ints(newResult[:])
						results[newResult] = true
					}
				}
			}
		}
	}

	slice := [][4]int{}
	for result, _ := range results {
		slice = append(slice, result)
	}
	return slice
}
