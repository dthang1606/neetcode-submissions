func minSubArrayLen(target int, nums []int) int {
	// move left if sum larger than target
	// output := len(nums)
	minArray := len(nums)+1
	left := 0 
	subSum := 0
	if subSum  == target {
		return 1
	}
	for right := 0 ; right < len(nums);right++ {
		subSum += nums[right]
		for subSum >= target{
            if right-left +1 < minArray  {
                minArray = right - left +1 
            }
            subSum = subSum - nums[left]
            left ++
        }
	}
    if minArray == len(nums)+1{
        return 0
    }
	return minArray
}
