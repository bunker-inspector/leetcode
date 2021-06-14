package problems

const Fail = 1

func IsInterleavingString(s1 string, s2 string, r string) bool {
	memos := make([][]int, len(s1))
	for i := 0; i < len(s1); i++ {
		memos[i] = make([]int, len(s2))
	}

	var isInterleaving func(int, int) bool
	isInterleaving = func(s1idx int, s2idx int) bool {
		if s1idx+s2idx > len(r) {
			memos[s1idx][s2idx] = Fail
			return false
		}
		if s1idx+s2idx == len(r) {
			return s1idx == len(s1) && s2idx == len(s2)
		}
		if s1idx < len(s1) && s2idx < len(s2) {
			memo := memos[s1idx][s2idx]

			if memo == Fail {
				return false
			}
		}
		if s1idx < len(s1) && r[s1idx+s2idx] == s1[s1idx] {
			if isInterleaving(s1idx+1, s2idx) {
				return true
			}
		}
		if s2idx < len(s2) && r[s1idx+s2idx] == s2[s2idx] {
			if isInterleaving(s1idx, s2idx+1) {
				return true
			}
		}
		if s1idx < len(s1) && s2idx < len(s2) {
			memos[s1idx][s2idx] = Fail
		}
		return false
	}

	return isInterleaving(0, 0)
}
