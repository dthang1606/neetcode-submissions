func twoSum(numbers []int, target int) []int {
	// 1,2,3,4 - target 3
	// output 1 2 
	// 1 3 5 6 8
	// 2 3
	low := 0
	right := len(numbers) - 1
	for low < right{
		computedNumber := numbers[low] + numbers[right]
		if computedNumber == target{
			return []int{low+1, right + 1 }
		}else if (computedNumber < target){
			low ++
		}else{
			right--
		}
	}
	return []int{low+1, right + 1 }
}

