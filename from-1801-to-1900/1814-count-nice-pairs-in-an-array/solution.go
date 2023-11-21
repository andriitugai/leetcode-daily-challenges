func countNicePairs(nums []int) int {
    reverse := func(a int) int {
        rev := 0
        for a > 0 {
            rev = rev * 10 + a % 10
            a /= 10
        }
        return rev
    }
    m := make(map[int]int)
    // nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
    // => nums[i] - rev(nums[i]) == nums[j] - rev(nums[j])
    // So keep track of the count of nums[idx] - rev(nums[idx])
    result := 0
    for _, num := range nums {
        key := num - reverse(num)
        result = (result + m[key]) % 1000000007
        m[key] += 1
    }
    return result
}