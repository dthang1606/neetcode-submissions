func productExceptSelf(nums []int) []int {
	// hashMap := make(map[int]int, len(nums))
	zeroCount :=0 
	output := []int{}
	sumExcludeZero := 1
	for _, v := range nums{
		if v == 0 {
			zeroCount++
		}else{
			sumExcludeZero = sumExcludeZero*v	
		}
	}

	for _, v := range nums{
		if zeroCount== 1{
			if v==0{
				output = append(output, sumExcludeZero)
			}else{
				output = append(output, 0)
			}
		}else if zeroCount > 1{
			output = append(output, 0)
		}else{
			output = append(output, sumExcludeZero/v)
		}
		
		
	}
	return output
}
// we can solve this problem with O(n^2) to inner scan array 
// for i, v: =range nums{
//    for ...
//}