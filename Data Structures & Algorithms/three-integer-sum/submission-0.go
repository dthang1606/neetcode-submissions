func threeSum(nums []int) [][]int {
	// -1 0 1 2 -1 -4
	// -4 -1 -1 0 1 2
	// 
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++{
		if i > 0 && nums[i] == nums[i-1] { continue }
		left, right := i+1, len(nums)-1

		for left < right{
			if nums[left]+nums[right]+nums[i]==0{
				result = append(result, []int{nums[i], nums[left], nums[right]})
				 // skip duplicates for left
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                // skip duplicates for right
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
				left++
				right--
			}else if (nums[left]+nums[right] +nums[i]<0){
				left ++
			}else{
				right -- 
			}
		}
	}
	return result
}
