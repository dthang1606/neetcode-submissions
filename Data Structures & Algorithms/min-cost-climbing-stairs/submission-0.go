func minCostClimbingStairs(cost []int) int {
    // if len(cost)<=2{
	// 	return 
	// }
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	// minCost := 0
	for i:=2; i< len(cost);i++{
		dp[i]=cost[i] + min(dp[i-1], dp[i-2])
	}
	return min(dp[len(dp)-1], dp[len(dp)-2])
}
