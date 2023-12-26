func incremovableSubarrayCount(nums []int) int {
    result := 0

    for i := 0; i < len(nums); i ++ {
        for j := i; j < len(nums); j ++ {
            prev := 0
            flag := 1
            for k := 0; k < len(nums); k ++ {
                if k < i || k > j {
                    if nums[k] > prev {
                        prev = nums[k]
                    } else {
                        flag = 0
                        break
                    }
                }
            }
            result += flag
        }
    }
    return result
}