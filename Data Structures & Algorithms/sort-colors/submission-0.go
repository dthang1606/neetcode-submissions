func sortColors(nums []int) {
    low := 0 
	mid := 0
	high := len(nums)-1
	//nums = [0,1,2,1,1,1]-> 0 0 1 1 2 2
	// 0 2 2 1 1 1 , 
	for mid <= high{
		if nums[mid] == 0 {
			// swap with low, mid =0 like : 2 0 -> 0 1 1 2
			nums[mid], nums[low] = nums[low], nums[mid]
			low ++
			mid ++
		}else if(nums[mid]==1){
			mid++
		}else{ // mid ==2
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}	
	}
}
