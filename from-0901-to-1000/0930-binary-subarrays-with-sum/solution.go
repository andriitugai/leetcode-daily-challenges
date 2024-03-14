func numSubarraysWithSum(nums []int, goal int) int {
    result := 0
    prefix := 0
    count := make(map[int]int)
    count[0] = 1

    for _, num := range nums {
        prefix += num
        result += count[prefix -  goal]
        count[prefix] += 1
    }
    return result
}