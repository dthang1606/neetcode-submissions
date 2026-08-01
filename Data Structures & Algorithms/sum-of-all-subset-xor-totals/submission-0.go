func subsetXORSum(nums []int) int {
	var res int
	var dfs func(int, int)
	dfs = func(i int, xor int) {
		if i == len(nums) {
			res += xor
			return
		}

		// take
		dfs(i+1, xor^nums[i])

		// skip
		dfs(i+1, xor)
	}

	dfs(0, 0)
	return res
}
