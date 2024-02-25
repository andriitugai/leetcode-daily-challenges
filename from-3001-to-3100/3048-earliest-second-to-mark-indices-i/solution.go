func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
    n := len(nums)
    m := len(changeIndices)

    check := func(val int) bool {
        marked := make([]int, n)
        for j := val - 1; j > -1; j -= 1 {
            i := changeIndices[j] - 1
            if j + 1 < val {
                val = j + 1
            }
            if marked[i] == 0 {
                val -= (nums[i] + 1)
                if val < 0 {
                    return false
                }
                marked[i] = 1
            }
        }
        for _, mark := range marked {
            if mark == 0 {
                return false
            }
        }
        return true
    }

    left, right := 0, m
    for left < right {
        mid := (left + right) / 2
        if check(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }

    if !check(right) {
        return -1
    }
    return right
}