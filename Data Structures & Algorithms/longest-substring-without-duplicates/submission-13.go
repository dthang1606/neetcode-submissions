func lengthOfLongestSubstring(s string) int {
	// zxyzczyc, xyzc, yzcz, czy
	result := 0
	left := 0
	lastSeen := make(map[byte]int, len(s))
	// lastSeen[s[left]] = left
	// abba
	// a b b a
	for right := 0; right < len(s); right ++{
		c := s[right]
		// jump left when ever we got the same char comapred with current
		if val, ok := lastSeen[c]; ok  && val >= left {
			left = val +1
		}

		if right - left + 1  > result{
			result = right - left +1
		}
		lastSeen[c] = right
		
	}
	return result
}
