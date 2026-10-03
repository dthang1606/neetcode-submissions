func topKFrequent(nums []int, k int) []int {
	//1 4
	//2 2
	//3 3
	//k = 2
	// for hashmap , if value > k , append into output
	// problem is  can't chose if simple rule
	// then need bucket sort,can use heap, ..
	hashmap := make(map[int]int, len(nums))
	bucket := make([][]int, len(nums)+1)
	for _,v := range nums{
		hashmap[v]++
	}
	// 1 2 2 3 3  , k=2
	// 1 1
	// 2 2
	// 3 2
	//bucket 1, 1; 2, 2 3; 3, 4 5
	for i, v := range hashmap{
		idx := v
		bucket[idx] = append(bucket[idx], i)
	}
	result := []int{}
	for i:= len(bucket)-1; i>=0 && len(result)<k; i--{
		result = append(result, bucket[i]...)

	}
	return result[:k]

	
}
