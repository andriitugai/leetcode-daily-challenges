/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

 type NestedIterator struct {
    ch chan int
    ans int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    ch := make(chan int)
    go Helper(nestedList,ch) // invoking the goroutine so that it can keep values in our channel, it will yield once val is kept in channel
    return &NestedIterator{ch,0}
}

func (this *NestedIterator) Next() int {
    return this.ans
}

func (this *NestedIterator) HasNext() bool {
    v,ok:=<-this.ch // ok will be true for non closed channels, reading from channel
    this.ans = v
    return ok 
}
func Helper(nestedList []*NestedInteger, ch chan int){
    var recurse func([]*NestedInteger)
    recurse = func(nestedList []*NestedInteger){
        for _,val := range nestedList{
            if val.IsInteger(){
                ch <- val.GetInteger() // writing to channel
            }else{
                recurse(val.GetList())
            }
        }
    }
    recurse(nestedList)
    defer close(ch) //closing the channels once all the values in nestedList are iterated
}