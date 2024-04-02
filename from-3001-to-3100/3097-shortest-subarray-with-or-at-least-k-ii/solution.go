func minimumSubarrayLength(nums []int, k int) int {
    bit_cnt := make([]int, 32)
    minLen := 1000000
    j, cur := 0, 0
    for i := 0; i < len(nums); i++ {
        cur |= nums[i]
        for b := 0; b < 32; b ++ {
            if nums[i] & (1 << b) > 0 {
                bit_cnt[b] += 1
            }
        }
        for cur >= k && j <= i {
            if i - j + 1 < minLen {
                minLen = i - j + 1
            }
            for b := 0; b < 32; b ++ {
                if nums[j] & (1 << b) > 0 {
                    bit_cnt[b] -= 1
                    if bit_cnt[b] == 0 {
                        cur ^= (1 << b)
                    }
                }
            }
            j += 1
        }
    }

    if minLen == 1000000 {
        return -1
    }
    return minLen
}