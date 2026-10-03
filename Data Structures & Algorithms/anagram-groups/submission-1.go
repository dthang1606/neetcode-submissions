func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)
	// idea is a =97 in asicii -> b =98, c=99, d=100 --> z =122
    for _, str := range strs {
        var count [26]int

        for _, c := range str {
			//meaning 26 chars and 'a'-'a' -> char a with index=0
            count[c-'a']++
        }

        groups[count] = append(groups[count], str)
    }

    result := make([][]string, 0, len(groups))

    for _, group := range groups {
        result = append(result, group)
    }

    return result
}