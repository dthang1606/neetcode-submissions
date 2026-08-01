func isInteger(s string) (int, error) {
    value, err := strconv.Atoi(s)
    return value, err
}
func compute(last int, prevLast int, opt string) int{
	var value int
	if opt == "+"{
		value = last+prevLast
	}else if opt =="-" {
		value = prevLast - last
	}else if opt =="*"{
		value = prevLast*last
	}else if opt == "/"{
		value = prevLast/last
	}
	return value
}
func evalRPN(tokens []string) int {
	var stack []int

	for _, v := range tokens{
		if value, err := isInteger(string(v));err==nil{
			stack = append(stack, value)
		}else{
			newVal := compute(stack[len(stack)-1], stack[len(stack)-2], string(v))
			// // pop 2 last
			stack = stack[:len(stack)-2]
			// push into stack
			stack = append(stack, newVal)
		}
	}
	return stack[len(stack)-1]
}
