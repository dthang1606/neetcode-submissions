func subsets(nums []int) [][]int {
	var result [][]int
	var dfs func (i int)
	var subset []int
	// hashmap := make(map[int]int, 0, 10)
	dfs = func (i int){
		if i== len(nums){
			tmp := make([]int, len(subset))
			copy(tmp, subset)
			result = append(result, tmp)
			return
		}
		// take 
		
		subset = append(subset, nums[i])
		dfs(i+1)
		// backtrack
		subset = subset[:len(subset)-1]
		// skip
		dfs(i+1)
	}
	dfs(0)
	return result
}
// need to take or exclude 
