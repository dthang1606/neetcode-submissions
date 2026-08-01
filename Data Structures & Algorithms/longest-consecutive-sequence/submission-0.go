func longestConsecutive(nums []int) int {
	// find quickly value -> hashmap problem
	hashMap := make(map[int]bool, len(nums))
	max := 0
	for _, v := range nums{
		hashMap[v] = true
	}
	// 4 20 2 10 3 4 5
	// only count if it is the start of a sequence
	// the start of a sequence is if nums[i]-1 doesn't exist 
	// tempArr := make(int[], len(nums))
	for _, v := range nums{
		// start from start of sequence
		if !hashMap[v-1]{
			// access here mean the start of sequece
			// start counting
			current := v
			length := 1
			for hashMap[current+1]{
				length++
				current++
			}

			if length > max{
				max = length
			}
		}

	}
	return max
}
