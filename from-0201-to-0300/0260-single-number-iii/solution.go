func singleNumber(nums []int) []int {
    xor := nums[0]
    for _, num := range nums[1:] {
        xor ^= num
    }
    mask := 1
    for mask & xor == 0 {
        mask <<= 1
    }

    var b1, b2 int = 0, 0
    for _, num := range nums {
        if mask & num != 0 {
            b1 ^= num
        } else {
            b2 ^= num
        }
    }
    return []int{b1, b2}
}