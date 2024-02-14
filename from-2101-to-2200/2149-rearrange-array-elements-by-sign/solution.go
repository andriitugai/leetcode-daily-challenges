func rearrangeArray(nums []int) []int {
    n := len(nums)
    pos, pp := make([]int, n/2), 0
    neg, np := make([]int, n/2), 0

    for _, num := range nums {
        if num > 0 {
            pos[pp] = num
            pp += 1
        } else {
            neg[np] = num
            np += 1
        }
    }

    pp, np = 0, 0
    for i := 0; i < n; i += 2 {
        nums[i] = pos[pp]
        pp += 1
        nums[i + 1] = neg[np]
        np += 1
    }
    return nums
}