func isPalindrome(s string) bool {
	var filtered []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filtered = append(filtered, unicode.ToLower(r))
		}
	}

	left := 0
	right := len(filtered) - 1

	for left < right {
		if filtered[left] == filtered[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}