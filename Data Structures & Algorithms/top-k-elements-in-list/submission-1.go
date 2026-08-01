func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int, len(nums))
	bucketStruct := make([][]int, len(nums)+1)
	// // count number of frequent elements into the hashmap
	for _, v:= range nums{
		hashMap[v] ++
	}

	// devide into the bucket
	for val, freq := range hashMap{
		idx := freq
		bucketStruct[idx] = append(bucketStruct[idx], val)
	}
	result := []int{}
	for i := len(bucketStruct)-1; i>0 && len(result)< k; i--{
		result = append(result, bucketStruct[i]...)
	}

	return result[:k]
}
