func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	result :=0

	left :=0
	right := len(people)-1
	for left <= right{
		sum := people[left]+people[right]
		if sum <= limit{
			left++
		}
		right --
		result++
	}
	// people=[3,5,3,4], 3 3 4 5
	// limit=5
	return result
}
