func isValid(s string) bool {
    var stack []rune
    for _, i := range s{
        if (i == '(') || (i == '[') || (i =='{'){
            stack = append(stack, i)
        }else if (len(stack)>0) && (i=='}' && stack[len(stack)-1] == '{'){
            stack = stack[:len(stack)-1]
        }else if (len(stack)>0) &&  (i==']' && stack[len(stack)-1] == '['){
            stack = stack[:len(stack)-1]
        }else if (len(stack)>0) &&  (i==')' && stack[len(stack)-1] == '('){
            stack = stack[:len(stack)-1]
        }else {
            // fmt.Println(stack)
            return false
        }
    }
    return len(stack)==0
}
