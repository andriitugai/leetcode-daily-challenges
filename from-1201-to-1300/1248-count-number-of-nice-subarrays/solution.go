func numberOfSubarrays(nums []int, k int) int {
    for i := 0; i < len(nums); i ++ {
        nums[i] %= 2
    }
    prefixSum := make([]int, len(nums) + 1)
    prefixSum[0] = 1

    result := 0
    count := 0

    for _, num := range nums {
        count += num
        if count >= k {
            result += prefixSum[count - k]
        }
        prefixSum[count] += 1
    }
    return result
}