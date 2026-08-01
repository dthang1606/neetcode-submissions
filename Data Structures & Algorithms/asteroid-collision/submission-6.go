func asteroidCollision(asteroids []int) []int {
	// stack, fmt.Println(math.Abs(x)) // Kết quả: 15.5
	stack := make([]int, 0, len(asteroids))
	//asteroids=[2,4,-4,-1]
	// stack = append(stack, asteroids[0])
	for i:=0; i< len(asteroids); i++{
        // fmt.Println(stack)
		for len(stack) > 0 && asteroids[i] <0 && stack[len(stack)-1] >0 &&AbsInt(asteroids[i]) > stack[len(stack)-1]{
			// pop last
            stack = stack[:len(stack)-1]
		}
        // need to add whenever last of stack compare
        if len(stack) > 0 {
            if asteroids[i] < 0 && AbsInt(asteroids[i]) < stack[len(stack)-1]{
                //skip i
            }else if (asteroids[i] < 0 && AbsInt(asteroids[i]) == stack[len(stack)-1]){
                // pop and skip i
                stack = stack[:len(stack)-1]
            }else{
                stack = append(stack, asteroids[i])
            }
        }else{
            stack = append(stack, asteroids[i])
        }
	}
	return stack
}

// Hàm tính giá trị tuyệt đối cho int
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
// 2 3 5 -2 4 -7 9
// 2 3 5 4 -7 9

// -7 9


