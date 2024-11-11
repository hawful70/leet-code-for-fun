type MinStack struct {
    mainStack []int
    minStack []int
}


func Constructor() MinStack {
    return MinStack {
        mainStack: []int{},
        minStack: []int{},
    }
}


func (this *MinStack) Push(val int)  {
    this.mainStack = append(this.mainStack, val)
    if len(this.minStack) == 0 {
        this.minStack = append(this.minStack, len(this.mainStack) - 1)
    } else {
        topIndexMin := this.minStack[len(this.minStack) - 1]
        if this.mainStack[topIndexMin] > val {
            this.minStack = append(this.minStack, len(this.mainStack) - 1)
        }
    }
}


func (this *MinStack) Pop()  {
    if len(this.mainStack) > 0 && len(this.minStack) > 0 {
        if this.minStack[len(this.minStack) - 1] == len(this.mainStack) - 1 {
            this.minStack = this.minStack[:len(this.minStack) - 1]
            this.mainStack = this.mainStack[:len(this.mainStack) - 1]
        } else {
            this.mainStack = this.mainStack[:len(this.mainStack) - 1]
        }
    }    
}


func (this *MinStack) Top() int {
    if len(this.mainStack) > 0 {
        return this.mainStack[len(this.mainStack) - 1]
    }
    return -1
}


func (this *MinStack) GetMin() int {
    index := this.minStack[len(this.minStack) - 1]
    return this.mainStack[index]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */