func maxSubArray(nums []int) int {
    left :=0
	sum :=nums[left]
	max :=sum
	//nums=[-2,1,-3,4,-1,2,1,-5,4]
	// 4 -1 2 1 = 6
	// 4 -1 2 1 -5 4 = 5 -> just compare "max" and "sum"
	for right :=1; right < len(nums); right ++{
		if nums[right] > sum && sum<0{
			left = right
			sum = nums[left]
			max = sum
		}else{
			sum +=nums[right]
			if max < sum{
				max = sum
			}
		}
	}
	return max
}
