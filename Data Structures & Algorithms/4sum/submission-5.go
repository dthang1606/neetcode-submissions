func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums) // nlogn
	//Input: nums = [3,2,3,-3,1,0], target = 3
	// -3 0 1 2 3 3
	// 
	n := len(nums)
	output := make([][]int, 0,n)
	for i := 0; i< n-3; i++{
		// skip duplicate
		if i > 0 && nums[i]==nums[i-1]{
			continue
		}
		for j:=i+1; j< n-2; j++{
			// skip duplicate
			if j > i+1 && nums[j]==nums[j-1]{
				continue
			}
			left := j+1
			right:= n-1
			for left < right {
				sum := nums[i]+nums[j]+nums[left]+nums[right]
				if sum == target{
					output = append(output, []int{nums[i],nums[j],nums[left],nums[right]}) 
					for left < right && nums[left] == nums[left +1]{
						left ++
					}

					for left < right && nums[right] == nums[right -1]{
						right --
					}
					left ++
					right -- 
				}else if (sum < target){
					left ++
				}else if(sum>target){
					right--
				}
			}
		}
	}
	return output
}
