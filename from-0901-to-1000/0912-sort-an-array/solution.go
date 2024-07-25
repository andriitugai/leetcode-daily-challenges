func sortArray(nums []int) []int {
    if len(nums) < 2 {
        return nums
    }
    mid := len(nums) / 2
    left := sortArray(nums[:mid])
    right := sortArray(nums[mid:])
    return merge(left, right)
}

func merge(left, right []int) []int {
    result := make([]int, len(left) + len(right))
    i, j := 0, 0
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result[i + j] = left[i]
            i += 1
        } else {
            result[i + j] = right[j]
            j += 1
        }
    }
    for i < len(left) {
        result[i + j] = left[i]
        i += 1
    }
    for j < len(right) {
        result[i + j] = right[j]
        j += 1
    }
    return result
}