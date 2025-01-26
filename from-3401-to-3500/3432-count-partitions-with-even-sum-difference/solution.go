func countPartitions(nums []int) int {
    total := 0
    for _, num := range(nums) {
        total += num
    }
    if total % 2 == 1 {
        return 0
    }

    return len(nums) - 1
}