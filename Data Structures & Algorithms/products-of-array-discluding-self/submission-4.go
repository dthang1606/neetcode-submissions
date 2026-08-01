func productExceptSelf(nums []int) []int {
	output := make([]int,0, len(nums))

	// leftPrefix := []int{1}
	previousSum := 1
	// Input: nums = [1,2,4,6] -> 1 1 2 8
	// build left prefix
	output = append(output, previousSum)
	for i:= 1; i< len(nums); i++{
		// previousSum = previousSum * nums[i-1]
		previousSum = previousSum*nums[i-1]
		output = append(output, previousSum)
	}
	// reset temp value to 1 default
	previousSum = 1
	// set the last element of nums, then we travel from len(nums)-2
	output[len(nums)-1] = 1* output[len(nums)-1]
	fmt.Println(output)
	// Input: nums = [1,2,4,6] -> 48 24 6 1
	// build right prefix
	for i := len(nums)-2; i>=0 ;i--{
		previousSum = previousSum*nums[i+1]
		output[i] = previousSum*output[i]
	}
	// output[0] = previousSum*nums[0]
	return output
}
// we can solve this problem with O(n^2) to inner scan array 
// for i, v: =range nums{
//    for ...
//}

// solve this problem with O(n) but maybe interviewer doesn't like it by using devision by 0
// func productExceptSelf(nums []int) []int {
// 	// hashMap := make(map[int]int, len(nums))
// 	zeroCount :=0 
// 	output := []int{}
// 	sumExcludeZero := 1
// 	for _, v := range nums{
// 		if v == 0 {
// 			zeroCount++
// 		}else{
// 			sumExcludeZero = sumExcludeZero*v	
// 		}
// 	}

// 	for _, v := range nums{
// 		if zeroCount== 1{
// 			if v==0{
// 				output = append(output, sumExcludeZero)
// 			}else{
// 				output = append(output, 0)
// 			}
// 		}else if zeroCount > 1{
// 			output = append(output, 0)
// 		}else{
// 			output = append(output, sumExcludeZero/v)
// 		}
		
		
// 	}
// 	return output
// }