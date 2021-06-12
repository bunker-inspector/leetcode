package problems

func LongestNonrepeatingSubstring(s string) int {
	var longest int
	var currLen int
	var currStartIndex int
	chars := make(map[rune]int)
	for idx, char := range s {
		seenIdx, seen := chars[char]
		if seen && seenIdx >= currStartIndex {
			chars[char] = idx
			if currLen > longest {
				longest = currLen
			}
			currLen = idx - seenIdx
			currStartIndex = seenIdx
		} else {
			currLen++
			chars[char] = idx
		}
	}
	if currLen > longest {
		longest = currLen
	}
	return longest
}
