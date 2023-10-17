type OrderedStream struct {
    capacity  int
    stream    []string
    ptr       int
}


func Constructor(n int) OrderedStream {
    return OrderedStream{capacity: n, stream: make([]string, n), ptr: 0}
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
    this.stream[idKey - 1] = value
    prevPtr := this.ptr

    for this.ptr < this.capacity && this.stream[this.ptr] != "" {
        this.ptr ++
    }
    return this.stream[prevPtr:this.ptr]
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */