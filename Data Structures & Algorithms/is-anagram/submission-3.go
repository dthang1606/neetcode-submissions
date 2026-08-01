func isAnagram(s string, t string) bool {
	hashS := make(map[rune]int, len(s))
	hashT := make(map[rune]int, len(t))
	output := true
	for _, v := range s {
		hashS[v] ++
	}
	for _, v := range t {
		hashT[v] ++
	}
	if len(s) != len(t){
		return false
	}
	// r:2 a:2 c:2 e:1
	for key, value := range hashT{
		if val, ok := hashS[key] ; ok {
			fmt.Printf("key: %c, val:%s \n", key, val)
			if val == value{
				continue
			}else{
				return false
			}
		}else{
			return false
		}
	}
	return output
}
