func rob(nums []int) int {
    if len(nums)<2{
        return nums[0]
    }
    if len(nums)<3{
        return max(nums[0], nums[1])
    }
    dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0],nums[1])

	for i:=2; i<len(nums); i++{
		dp[i]= max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[len(dp)-1]
}
// greedy mindset: 2 9 8 3 6 -> "What looks best NOW?"
// -> output: 9+3=12
// dynamic programing mindset --> "What is the best result UP TO NOW?"
// sub problem : 2 9 10 12 16
// -> output: max(2+8+6, 3+9)