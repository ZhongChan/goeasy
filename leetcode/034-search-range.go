package leetcode

func searchRange(nums []int, target int) []int {
	left := searchBoarder(nums, target, true)
	right := searchBoarder(nums, target, false)
	return []int{left, right}
}

func searchBoarder(nums []int, target int, flag bool) int {
	boarder := -1
	l := 0
	r := len(nums) - 1

	//闭区间 [l,r]
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			//找到target了
			if flag {
				//左边界 - 区间收缩到左边，不会再找到 target了
				r = mid - 1
			} else {
				//右边界 - 区间收缩，有可能还能找到 target
				l = mid + 1
			}

			boarder = mid
		}
	}

	return boarder
}
