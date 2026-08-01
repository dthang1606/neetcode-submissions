func subsetXORSum(nums []int) int {
	// O(n)
	or := 0
	for _, num := range nums {
		or |= num
	}
	return or * (1 << (len(nums) - 1))
}

// brute force: 
// var res int
// var dfs func(i int, xor int)
// dfs = func(i int, xor int){
// 	if i == len(nums){
// 		res += xor
// 		return
// 	}

// 	// take
// 	dfs(i+1, xor^nums[i])
// 	// skip
// 	dfs(i+1, xor)
// }
// dfs(0,0)
// return res
// or := 0
// for _, num := range nums {
// 	or |= num
// }
// return or * (1 << (len(nums) - 1))
