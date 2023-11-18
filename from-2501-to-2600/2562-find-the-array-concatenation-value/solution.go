func findTheArrayConcVal(nums []int) int64 {
    concat := func(a, b int) int {
        if b < 10 {
            return a * 10 + b
        } else if b < 100 {
            return a * 100 + b
        } else if b < 1000 {
            return a * 1000 + b
        } else if b < 10000 {
            return a * 10000 + b
        }
        return a * 100000 + b
    }
    var result int64 = 0
    left, right := 0, len(nums) - 1
    for left < right {
        result += int64(concat(nums[left], nums[right]))
        left += 1
        right -= 1
    }
    if left == right {
        result += int64(nums[left])
    }
    return result
}