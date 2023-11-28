func sumIndicesWithKSetBits(nums []int, k int) int {
    countBits := func(n int) int {
        c := 0
        for n > 0 {
            n = n & (n - 1)
            c += 1
        }
        return c
    }
    result := 0
    for i, num := range nums {
        if countBits(i) == k {
            result += num
        }
    }
    return result
}