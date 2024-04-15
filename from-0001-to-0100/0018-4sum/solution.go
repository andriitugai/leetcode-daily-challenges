func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    n := len(nums)

    for i := 0; i < n; i ++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        for j := i + 1; j < n; j ++ {
            if j > i + 1 && nums[j] == nums[j - 1] {
                continue
            }
            left := j + 1
            right := n - 1
            for left < right {
                sum := nums[i] + nums[j] + nums[left] + nums[right]
                if sum == target {
                    result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
                    low := nums[left]
                    high := nums[right]
                    left += 1
                    right -= 1
                    for left < right && nums[left] == low {
                        left += 1
                    }
                    for left < right && nums[right] == high {
                        right -= 1
                    }
                } else if sum > target {
                    right -= 1
                } else {
                    left += 1
                }
            }
        }
    }
    return result
}