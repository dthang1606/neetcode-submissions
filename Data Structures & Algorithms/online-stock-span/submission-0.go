type StockSpanner struct {
	prices []int
	stack []int
}

func Constructor() StockSpanner {
	return StockSpanner {
		prices: make([]int, 0),
		stack: make([]int, 0),
	}
}
// 100 80 60 70 60 75 85
// 1   1  1   2  1 4 6
func (this *StockSpanner) Next(price int) int {
	output :=1
	this.prices = append(this.prices, price)
	i := len(this.prices)-1
	this.stack = append(this.stack, i)
	// temp := []int{}
	for len(this.stack)> 0 &&  this.prices[i] >= this.prices[this.stack[len(this.stack)-1]]{
		//temp stack was removed
		// temp = append(temp, this.stack[:len(this.stack)-1])
		// output = i - this.stack[len(this.stack)-1]+1
		// pop last to travel back
		this.stack = this.stack[:len(this.stack)-1]
	}
	if len(this.stack)==0{
		output = i+1
	}else{
		output = i - this.stack[len(this.stack)-1]
	}
	this.stack = append(this.stack, i)
	return output
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */
 