func rotate(nums []int, k int) {
	// way 1 use extra space
	n := len(nums)
	k = k % n
	result := make([]int, n)
	// new index would be: newIndex = (i+k)%n
	for i := 0 ; i< n ; i++{
		result[(i+k)%n] = nums[i]
	} 
	copy(nums, result )
}


// left :=0 
// right:= k

// for left < k && right<len(nums){
// 	nums[left], nums[right] = nums[right], nums[left]
// 	right++
// 	left++
// }