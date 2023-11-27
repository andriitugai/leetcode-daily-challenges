func maxSum(nums []int) int {
    getMaxDigit := func(num int) int {
        md := -1
        for num > 0 {
            if num % 10 > md {
                md = num % 10
            }
            num /= 10
        }
        return md
    }
    maxSum := -1
    for i := 0; i < len(nums) - 1; i ++ {
        for j := i + 1; j < len(nums); j ++ {
            if getMaxDigit(nums[i]) == getMaxDigit(nums[j]) {
                sum := nums[i] + nums[j]
                if sum > maxSum {
                    maxSum = sum
                }
            }
        }
    }
    return maxSum
}