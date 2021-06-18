package problems

import (
	"strings"
)

/*
   P   A   H   N
   A P L S I I G
   Y   I   R
   PAHNAPLSIIGYIR

   P     I    N
   A   L S  I G
   Y A   H R
   P     I
   PINALSIGYARHPI

   P     I     N     R
   A   L S   I G   F I
   Y A   H R   F R   E D
   P     I     O     N

   PAYPALISHIRINGFORFRIEND

   P       H       R
   A     S I     O F
   Y   I   R   F   R
   P L     I G     I D
   A       N       N

   PAYPALISHIRINGFORFRIEND

   Val 1   : i * (r-1) + i
   Not End : i * (r-1) + 2 * (r - i)
*/

func ConvertZigZag(s string, rows int) string {
	if s == "" || rows == 0 {
		return ""
	}
	if rows == 1 {
		return s
	}

	var result strings.Builder
	var i int
	// Cleaner to handle 1st and last rows separately since they
	// only have characters at the peaks / troughs
	// `(2 * rows) - 2` is the distance between the first
	// character in every zigzag "loop", i.e. any 2 peaks / troughs
	for i = 0; i < len(s); i += (2 * rows) - 2 {
		result.WriteByte(s[i])
	}
	// Handle inner rows
	for row := 1; row < rows-1; row++ {
		i := row
		for i < len(s) {
			result.WriteByte(s[i])
			// The distance in the original string between the first character
			// in this loop and the second
			secondIdx := i + 2*rows - 2*(row+1)
			if secondIdx < len(s) {
				result.WriteByte(s[secondIdx])
			}
			i += (2 * rows) - 2
		}
	}
	for i = rows - 1; i < len(s); i += (2 * rows) - 2 {
		result.WriteByte(s[i])
	}
	return result.String()
}
