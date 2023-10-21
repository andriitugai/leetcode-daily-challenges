func constrainedSubsetSum(nums []int, k int) int {
    if len(nums) == 0 {
        return 0
    }
    
    var dq deque
    r := nums[0]
    
    for i := range nums {
        if dq.Left() != -1 {
            nums[i] += dq.Left()
        }
        
        r = max(r, nums[i])
        
        for dq.Right() != -1 && nums[i] > dq.Right() {
            dq.PopRight()
        }
        
        if nums[i] > 0 {
            dq.PushRight(nums[i])
        }
        
        if i - k >= 0 && dq.Left() != -1 && dq.Left() == nums[i - k] {
            dq.PopLeft()
        }
    }
    
    return r
}

func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}

type deque []int

func (dq *deque) Left() int {
    if len(*dq) == 0 {
        return -1
    }
    
    return (*dq)[0]
}

func (dq *deque) PopLeft() {
    if len(*dq) == 0 {
        return
    }
    
    *dq = (*dq)[1:]
}

func (dq *deque) Right() int {
    if len(*dq) == 0 {
        return -1
    }
    
    return (*dq)[len(*dq)-1]
}

func (dq *deque) PopRight() {
    if len(*dq) == 0 {
        return
    }
    
    *dq = (*dq)[:len(*dq)-1]
}

func (dq *deque) PushRight(v int) {
    *dq = append(*dq, v)
}