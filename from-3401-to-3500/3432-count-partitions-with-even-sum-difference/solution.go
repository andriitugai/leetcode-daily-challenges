func countPartitions(nums []int) int {
    n := len(nums)
    total := 0
    for _, num := range(nums) {
        total += num
    }
    if total % 2 == 1 {
        return 0
    }

    curr := 0
    result := 0
    for i := 0; i < n - 1; i ++ {
        curr += nums[i]
        total -= nums[i]
        if total % 2 == curr % 2 {
            result += 1
        }
    }
    return result
}