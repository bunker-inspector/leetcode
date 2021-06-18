package problems

type startRunPair struct {
	Start int
	Run   int
}

func isPalindromeFrom(s *string, start int, run int, cache *map[startRunPair]bool) bool {
	res, cached := (*cache)[startRunPair{Start: start, Run: run}]
	if cached {
		return res
	}
	if run == 1 {
		return true
	}
	if run == 2 {
		return (*s)[start] == (*s)[start+1]
	}

	result := (*s)[start] == (*s)[start+run-1] && isPalindromeFrom(s, start+1, run-2, cache)
	(*cache)[startRunPair{Start: start, Run: run}] = result

	return result
}

func LongestPalindrome(s string) string {
	cache := make(map[startRunPair]bool)

	for l := len(s); l > 0; l-- {
		for i := 0; i+l <= len(s); i++ {
			if isPalindromeFrom(&s, i, l, &cache) {
				return s[i : i+l]
			}
		}
	}
	return ""
}
