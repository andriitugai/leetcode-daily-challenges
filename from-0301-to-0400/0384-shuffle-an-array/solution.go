type Solution struct {
    orig []int
    curr []int
}


func Constructor(nums []int) Solution {
    curr := make([]int, len(nums))
    copy(curr, nums)
    return Solution{orig: nums, curr: curr}
}


func (this *Solution) Reset() []int {
    copy(this.curr, this.orig)
    return this.curr
}


func (this *Solution) Shuffle() []int {
    rand.Shuffle(len(this.curr), func(i, j int){
        this.curr[i], this.curr[j] = this.curr[j], this.curr[i]
    })
    return this.curr
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */