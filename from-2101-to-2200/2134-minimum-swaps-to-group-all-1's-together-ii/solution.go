func minSwaps(nums []int) int {
    n := len(nums)
    nOnes := 0
    for i := 0; i < n; i ++ {
        if nums[i] == 1 {
            nOnes += 1
        }
    }
    if nOnes < 2 {
        return 0
    }

    curOnes := 0
    for i := 0; i < nOnes; i ++ {
        if nums[i] == 1 {
            curOnes += 1
        }
    }
    maxOnes := curOnes
    for i := nOnes; i < n + nOnes; i ++ {
        curOnes += nums[i % n] - nums[(i - nOnes + n) % n]
        if curOnes > maxOnes {
            maxOnes = curOnes
        }
    }

    return nOnes - maxOnes
}