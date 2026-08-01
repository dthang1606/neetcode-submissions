func rotate(nums []int, k int) {
	n:=len(nums)
	k = k%n
	// reverse all array
	reverse(nums, 0,n-1)
	// reverse 0-k
	reverse(nums, 0, k-1)
	// reverse the rest
	reverse(nums, k, n-1) 
}

func reverse(nums []int, left int, right int){
	for left < right{
		nums[left], nums[right] = nums[right], nums[left]
		right--
		left++
	}
}
// func rotate(nums []int, k int) {
// 	// way 1 use extra space
// 	n := len(nums)
// 	k = k % n
// 	result := make([]int, n)
// 	// new index would be: newIndex = (i+k)%n
// 	for i := 0 ; i< n ; i++{
// 		result[(i+k)%n] = nums[i]
// 	} 
// 	copy(nums, result )
// }


// left :=0 
// right:= k

// for left < k && right<len(nums){
// 	nums[left], nums[right] = nums[right], nums[left]
// 	right++
// 	left++
// }