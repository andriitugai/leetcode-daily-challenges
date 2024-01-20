func sumSubarrayMins(arr []int) int {
	// this forces everything out of the stack at the end
    arr = append(arr, -1)
    min := NewStack()
    sum := 0
    for i := range arr {
        for min.Len() > 0 && arr[i] < arr[min.Top()] {
            //[p.....][c][......i]
			//    P           K
			// the number of subarrays with c as the minimum is P * K i.e. (c-p) * (i-c) see bottom
            c, p := min.Pop(), min.Top()
			P, K := c - p, i - c
            sum += arr[c] * P * K         // see proof below
        }
        
        min.Push(i)
    }
    
    return sum % int(1e9+7)
}

type stack struct {
    Push    func(x int)
    Pop     func() int
    Top     func() int
    Len     func()  int
}

func NewStack() stack {
    arr := make([]int, 0)
    return stack {
        Push: func(x int) {
            arr = append(arr, x)
        },
        Pop: func() int {
            item := arr[len(arr)-1]
            arr = arr[:len(arr)-1]
            return item
        },
        Top: func() int {
            if len(arr) == 0 {
                return -1
            }
            return arr[len(arr)-1]
        },
        Len: func() int {
            return len(arr)
        },
    }
}