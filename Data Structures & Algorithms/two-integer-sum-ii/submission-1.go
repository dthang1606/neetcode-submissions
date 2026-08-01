func twoSum(numbers []int, target int) []int {
	// var result []int
	right := len(numbers)-1
	left := 0 
	for left < right {
		if numbers[left] + numbers[right] < target{
			left ++ 
		}else if  numbers[left] + numbers[right] > target{
			right --
		}else{
			return []int{left + 1, right + 1}

		}
	}
	return []int{left + 1, right + 1}
}