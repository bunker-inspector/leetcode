package problems

import (
	"math"
	"strconv"
	"strings"
)

func ReverseInteger(x int) int {
	var builder strings.Builder
	var negative bool
	if x < 0 {
		negative = true
		x = -x
	}
	s := strconv.Itoa(x)
	for i := len(s) - 1; i >= 0; i-- {
		builder.WriteByte(s[i])
	}
	res, _ := strconv.Atoi(builder.String())
	if negative {
		res = -res
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}
