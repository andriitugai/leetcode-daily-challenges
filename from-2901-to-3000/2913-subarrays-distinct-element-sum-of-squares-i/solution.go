func sumCounts(nums []int) int {
    result := 0
    for i := 0; i < len(nums); i++ {
        s := make(map[int]bool)
        k := 0
        for _, num := range nums[i:] {
            if !s[num] {
                s[num] = true
                k += 1
            }
            result += k * k
        }
    }
    return result
}