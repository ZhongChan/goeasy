package leetcode

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	//闭区间[l,r]
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			//左侧区间
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r + 1
}
