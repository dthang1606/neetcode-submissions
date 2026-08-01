func lengthOfLongestSubstring(s string) int {
	left := 0
	lastDuplicateTrack := make(map[byte]int,len(s))
	output := 0 
	//abba
	for right :=0 ; right < len(s); right++{
		// invalidate condition, then move left
		if val, ok := lastDuplicateTrack[s[right]]; ok && val >=left {
			left = val + 1 
		}
		// add tracking sub array, for move left purpose
		lastDuplicateTrack[s[right]] = right
		// calculate max 
		temp := right - left + 1
		if temp > output {
			output = temp
		}
	}
	return output
}
