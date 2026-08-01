type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{

	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0{
		this.minStack = append(this.minStack, val)
	}else if (this.GetMin() > val){
		this.minStack = append(this.minStack, val)
	}else{
		this.minStack = append(this.minStack, this.GetMin())
	}
	// 2 5 1 4
	// stack: 2 5
	// min stack: 2 2 
	this.stack = append(this.stack, val)
	// fmt.Println(this.stack)
	// fmt.Println(this.minStack)

	
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
