func searchRange(nums []int, target int) []int {
    binSearch := func(isLeftmost bool) int {
        l, r := 0, len(nums) - 1
        res := -1

        for l <= r {
            m := l + (r - l) / 2
            if nums[m] > target {
                r = m - 1
            } else if nums[m] < target {
                l = m + 1
            } else {
                res = m
                if isLeftmost {
                    r = m - 1
                } else {
                    l = m + 1
                }
            }
        }
        return res
    }
    return []int{binSearch(true), binSearch(false)}
}
