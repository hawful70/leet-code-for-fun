type MyStack struct {
    stacks []int
}


func Constructor() MyStack {
    return MyStack{
        stacks: []int{},
    }
}


func (this *MyStack) Push(x int)  {
    this.stacks = append(this.stacks, x)
}


func (this *MyStack) Pop() int {
    top := this.stacks[len(this.stacks) - 1]
    this.stacks = this.stacks[:len(this.stacks) - 1]
    return top
}


func (this *MyStack) Top() int {
    return this.stacks[len(this.stacks) - 1]
}


func (this *MyStack) Empty() bool {
    if len(this.stacks) == 0 {
        return true
    }
    return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */