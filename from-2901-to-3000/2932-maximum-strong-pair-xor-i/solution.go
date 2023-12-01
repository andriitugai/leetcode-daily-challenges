func maximumStrongPairXor(nums []int) int {
    isStrongPair := func(a, b int) bool {
        v1 := a - b
        if v1 < 0 {
            v1 = -v1
        }
        v2 := a
        if b < a {
            v2 = b
        }
        return v1 <= v2
    }
    maxXor := -1
    for i := 0; i < len(nums); i ++ {
        for j := i; j < len(nums); j ++ {
            if isStrongPair(nums[i], nums[j]) {
                xor := nums[i] ^ nums[j]
                if xor > maxXor {
                    maxXor = xor
                }
            }
        }
    }
    return maxXor
}