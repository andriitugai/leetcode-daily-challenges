func minOperations(nums []int) int {
    n := len(nums)
    numOnes := 0
    for _, num := range nums {
        if num == 1 {
            numOnes += 1
        }
    }
    if numOnes > 0 {
        return n - numOnes
    }

    res := math.MaxInt32
    for i := 0; i < n; i ++ {
        g := nums[i]
        for j := i + 1; j < n; j ++ {
            g = gcd(g, nums[j])
            if g == 1 {
                res = min(res, j - i + n - 1)
                break
            }
        }
        if g != 1 {
            break
        }
    }
    if res == math.MaxInt32 {
        return -1
    }
    return res
}

func gcd(a, b int) int {
    if a == 0 {
        return b
    }
    return gcd(b % a, a)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}