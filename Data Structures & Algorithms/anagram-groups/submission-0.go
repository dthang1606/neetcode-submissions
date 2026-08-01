func groupAnagrams(strs []string) [][]string {
	sortedStringHashMap := make(map[string][]string, len(strs))
	result := make([][]string,0,len(strs))
	for _, str := range strs{
		oldString := str
		runStr := []rune(str)
		sort.Slice(runStr, func(i, j int) bool {
			return runStr[i] < runStr[j]
		})
		sortedStringHashMap[string(runStr)] = append(sortedStringHashMap[string(runStr)], oldString)
	}

	for _, v := range sortedStringHashMap{
		result = append(result, v)
	}
	return result
}
