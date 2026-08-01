func characterReplacement(s string, k int) int {
	// store English char counter
	freq := make(map[byte]int, len(s))
	result := 0
	left := 0 
	maxFreq := 0
	for right :=0; right < len(s); right++{
		freq[s[right]]++
		// find the max Freq char
		if freq[s[right]]>maxFreq{
			maxFreq = freq[s[right]]
		}
		// invalid window
		if ((right -left +1) - maxFreq > k) {
			freq[s[left]]--
			left ++
		} 

		if right -left +1 > result {
			result = right -left +1
		}
	}
	return result
}
