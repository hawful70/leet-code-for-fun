type MyQueue struct {
    stackOne []int
    stackTwo []int
}


func Constructor() MyQueue {
    return MyQueue {
        stackOne: []int{},
        stackTwo: []int{},
    }
}


func (this *MyQueue) Push(x int)  {
    this.stackOne = append(this.stackOne, x)
}


func (this *MyQueue) Pop() int {
    if len(this.stackTwo) == 0 {
        for i := len(this.stackOne) - 1; i >= 0; i-- {
            this.stackTwo = append(this.stackTwo, this.stackOne[i])
            this.stackOne = this.stackOne[:len(this.stackOne) - 1]
        }
    }
    topQueue := this.stackTwo[len(this.stackTwo) - 1]
    this.stackTwo = this.stackTwo[:len(this.stackTwo) - 1]
    return topQueue
}


func (this *MyQueue) Peek() int {
    topQueue := 0
    if len(this.stackTwo) > 0 {
        topQueue = this.stackTwo[len(this.stackTwo) - 1]
        return topQueue
    } else {
        for i := len(this.stackOne) - 1; i >= 0; i-- {
            this.stackTwo = append(this.stackTwo, this.stackOne[i])
            this.stackOne = this.stackOne[:len(this.stackOne) - 1]
        }
        topQueue = this.stackTwo[len(this.stackTwo) - 1]
        return topQueue
    }
    return -1
}


func (this *MyQueue) Empty() bool {
    return len(this.stackOne) == 0 && len(this.stackTwo) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */