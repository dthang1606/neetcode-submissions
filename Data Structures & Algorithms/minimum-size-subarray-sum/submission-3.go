func minSubArrayLen(target int, nums []int) int {
	minArr := len(nums)+1
	sum :=0
	left :=0 
	for right :=0 ; right <len(nums); right ++ {
		sum += nums[right]
		for sum >= target {
			if right - left +1 < minArr{
				minArr = right - left +1
			}

			sum -=nums[left]
			left++
		} 
	}
	if minArr == len(nums)+1{
		return 0
	}
	return minArr
}
