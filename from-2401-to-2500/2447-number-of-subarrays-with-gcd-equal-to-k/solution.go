func subarrayGCD(nums []int, k int) int {
    gcd := func(a, b int) int {
        for b > 0 {
            a, b = b, a % b
        }
        return a
    }
    result := 0
    for i := 0; i < len(nums); i ++ {
        m := nums[i]
        for j := i; j < len(nums); j ++ {
            m = gcd(m, nums[j])
            if m == k {
                result += 1
            } else if m < k {
                break
            }
        }
    }
    return result
}