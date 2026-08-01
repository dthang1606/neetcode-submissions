type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	// hello world
	// 4#hello4#world
	result := ""
	for _, v := range strs{
		result += fmt.Sprintf("%d#%s", len(v), v)
	}
	return result
}

func (s *Solution) Decode(encoded string) []string {
	// 4#hello4#world
	result := []string{}
	i:=0
	for i < len(encoded){
		j := i
		// find #
		for encoded[j] != '#'{
			j++
		}

		// get length
		length, _ := strconv.Atoi(encoded[i:j])
		start := j+1
		end := start + length

		result = append(result, encoded[start:end])
		i = end
	}
	return result
}
