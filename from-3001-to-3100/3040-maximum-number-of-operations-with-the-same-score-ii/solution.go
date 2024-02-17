func max3(v1, v2, v3 int) int {
    if v1 > v2 {
        if v3 > v1 {
            return v3
        } else {
            return v1
        }
    } else {
        if v3 > v2 {
            return v3
        }
    }
    return v2
}

type state struct {
    left int
    right int
    s int
}

func maxOperations(nums []int) int {
    memo := make(map[state]int)

    var getMaxOps func(left, right, s int) int
    getMaxOps = func(left, right, s int) int {
        if right - left < 1 {
            return 0
        }
        curState := state{left: left, right: right, s: s}
        if val, ok := memo[curState]; ok {
            return val
        }
        v1 := 0
        if nums[left] + nums[left + 1] == s {
            v1 = 1 + getMaxOps(left + 2, right, s)
        }
        v2 := 0
        if nums[left] + nums[right] == s {
            v2 = 1 + getMaxOps(left + 1, right - 1, s)
        }
        v3 := 0
        if nums[right - 1] + nums[right] == s {
            v3 = 1 + getMaxOps(left, right - 2, s)
        }
        memo[curState] = max3(v1, v2, v3)
        return memo[curState]
    }

    n := len(nums)
    op1 := getMaxOps(0, n - 1, nums[0] + nums[1])
    op2 := getMaxOps(0, n - 1, nums[0] + nums[n - 1])
    op3 := getMaxOps(0, n - 1, nums[n-2] + nums[n - 1])

    return max3(op1, op2, op3)
}