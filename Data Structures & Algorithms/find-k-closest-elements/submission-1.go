func findClosestElements(arr []int, k int, x int) []int {
	left :=0
	right :=len(arr)-k
	//  [2,4,5,8,9 10],k=3, x=7, [5,8,9], [8 9 10]
	for left < right{
		mid := left + (right-left)/2

		if x-arr[mid] > arr[mid+k]-x{
			left = mid +1
		}else{
			right = mid
		}

	}
	return arr[left: left+k]
}
