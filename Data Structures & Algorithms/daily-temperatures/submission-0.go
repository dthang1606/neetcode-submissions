func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))  // store number of days waiting
	stack := make([]int, len(temperatures)) // store index

	for i:=0; i< len(temperatures); i++{
		for len(stack) > 0 && temperatures[i]>temperatures[stack[len(stack)-1]]{
			previousIndex := stack[len(stack)-1]
			// pop last //30,38,30,36,35,40,28---- 38 30 36
			stack = stack[:len(stack)-1]
			result[previousIndex] = i - previousIndex
		}
		stack = append(stack, i)
	}
	return result
}
