package problems

const (
	Left    = '('
	Right   = ')'
	Freebie = '*'
)

func ValidParenthesis(input string) bool {
	var p int
	for _, c := range input {
		if c == Left || c == Freebie {
			p++
		} else {
			p--
		}
		if p < 0 {
			return false
		}
	}
	p = 0
	for i := len(input) - 1; i >= 0; i-- {
		c := input[i]
		if c == Right || c == Freebie {
			p++
		} else {
			p--
		}
		if p < 0 {
			return false
		}
	}
	return true
}
