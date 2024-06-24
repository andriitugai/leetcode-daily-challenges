func minKBitFlips(nums []int, k int) int {
    flips := 0
    flipped := 0
    for i := 0; i < len(nums); i += 1 {
        if i >= k && nums[i - k] < 0 {
            flipped = (flipped + 1) % 2
        }

        if nums[i] == flipped {
            if i > len(nums) - k {
                return -1
            }

            flipped = (flipped + 1) % 2
            flips += 1
            nums[i] = -flips
        }
    }

    return flips
}