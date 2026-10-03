func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int, len(nums))

	for i, v := range nums{
		comparedNumber := target -v 
		if val, ok := hashMap[comparedNumber]; ok{
			return []int{val, i }
		}
		hashMap[v]=i
	}
	return []int{}
}

