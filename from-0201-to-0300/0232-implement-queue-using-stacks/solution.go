type MyQueue struct {
    S1 []int
    S2 []int
}


func Constructor() MyQueue {
    return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
    this.S1 = append(this.S1, x)
}


func (this *MyQueue) Pop() int {
    if len(this.S2) == 0 {
        for len(this.S1) != 0 {
            index := len(this.S1) - 1
            element := this.S1[index]
            this.S1 = this.S1[:index]
            this.S2 = append(this.S2, element)
        }
    }
    
    index := len(this.S2) - 1
    element := this.S2[index]
    this.S2 = this.S2[:index]
    
    return element
}


func (this *MyQueue) Peek() int {
    if len(this.S2) == 0 {
        for len(this.S1) != 0 {
            index := len(this.S1) - 1
            element := this.S1[index]
            this.S1 = this.S1[:index]
            this.S2 = append(this.S2, element)
        }
    }
    
    index := len(this.S2) - 1
    element := this.S2[index]
    
    return element
}


func (this *MyQueue) Empty() bool {
    if len(this.S1) == 0 && len(this.S2) == 0 {
        return true
    }
    return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */