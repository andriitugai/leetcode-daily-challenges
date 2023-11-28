func minOperations(nums []int, k int) int {
    visited := make(map[int]bool)
    ops := 0
    cnt := 0
    for i := len(nums) - 1; i >= 0 && cnt < k; i -- {
        if nums[i] <= k && !visited[nums[i]] {
            cnt += 1
            visited[nums[i]] = true
        }
        ops += 1
    }
    return ops
}