package problems

func findRotation(lidx, ridx int, nums *[]int) int {
	mid := (ridx + lidx) / 2
	if (*nums)[mid-1] > (*nums)[mid] {
		return mid
	}
	if (*nums)[lidx] > (*nums)[mid] {
		return findRotation(lidx, mid, nums)
	}
	return findRotation(mid, ridx, nums)
}

func findInRotated(lidx, ridx, target, rot int, nums *[]int) int {
	mid := (ridx + lidx) / 2
	ck := (mid + rot) % len(*nums)
	if (*nums)[ck] == target {
		return ck
	}
	if mid == lidx {
		return -1
	}
	if (*nums)[ck] > target {
		return findInRotated(lidx, mid, target, rot, nums)
	}
	return findInRotated(mid, ridx, target, rot, nums)
}

func FindInRotatedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	var rotation int
	if nums[0] > nums[len(nums)-1] {
		rotation = findRotation(0, len(nums), &nums)
	}
	return findInRotated(0, len(nums), target, rotation, &nums)
}
