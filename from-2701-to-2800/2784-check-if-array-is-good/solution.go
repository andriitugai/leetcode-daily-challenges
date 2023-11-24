func isGood(nums []int) bool {
    cnt := make(map[int]int)
    n := len(nums) - 1
    for i := 0; i <= n; i ++ {
        cnt[nums[i]] += 1
    }

    for i := 1; i < n; i ++ {
        if cnt[i] != 1 {
            return false
        }
    }

    return cnt[n] == 2
}