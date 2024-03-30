func subarraysWithKDistinct(nums []int, k int) int {
    var window func(kk int) int 
    window = func(kk int) int {
        left := 0
        result := 0
        contains := make(map[int]bool)
        counts := make(map[int]int)
        for right, num := range nums {
            contains[num] = true
            counts[num] += 1

            for len(contains) > kk {
                counts[nums[left]] -= 1
                if counts[nums[left]] == 0 {
                    delete(contains, nums[left])
                }
                left += 1
            }
            result += (right - left + 1)
        }
        return result
    }
    return window(k) - window(k - 1)
}