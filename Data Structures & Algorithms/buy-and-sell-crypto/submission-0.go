func maxProfit(prices []int) int {
	totalDays := len(prices)
	maxProfitPrice := 0
	left := 0
	right := left + 1
	//prices=[7,1,5,3,6,4]
	for left < totalDays && right < totalDays{
		tempProfit := prices[right] - prices[left]
		if tempProfit > maxProfitPrice {
			maxProfitPrice = tempProfit
		}
		if prices[left] > prices[right] {
			left = right
			right++
		}else{
			right++
		}
		
	}
	return maxProfitPrice
}
