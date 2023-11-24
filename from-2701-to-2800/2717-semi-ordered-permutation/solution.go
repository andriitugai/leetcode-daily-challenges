func semiOrderedPermutation(nums []int) int {
    p1, pn := -1, -1
    n := len(nums)

    for i := 0; i < n; i ++ {
        if nums[i] == 1 {
            p1 = i
        } else if nums[i] == n {
            pn = i
        }
    }
    result := p1 + n - pn - 1
    if pn < p1 {
        result -= 1
    }
    return result
}