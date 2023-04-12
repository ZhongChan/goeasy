package leetcode

func search(nums []int, target int) int {
	// [left,right]
	// [-1,0,1,3,9,12] 9
	right := len(nums) - 1
	left := 0
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			//比区间 mid 已经比较过了，需要 -1
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func search2(nums []int, target int) int {
	//左闭右开 [left,right)
	left := 0
	right := len(nums)

	//区间最右边不会被选中
	for left < right {
		//防止数据溢出
		mid := right + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			//不需要再和mid 比较了，所以需要 + 1
			left = mid + 1
		} else {
			// [left,right) 开区间 mid 不会被比较
			right = mid
		}
	}

	return -1
}
