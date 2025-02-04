func maxAscendingSum(nums []int) int {
    curSum, maxSum := 0, 0
    prev := 0
    for _, curr := range nums {
        if curr > prev {
            curSum += curr
        } else {
            if curSum > maxSum {
                maxSum = curSum
            }
            curSum = curr
        }
        prev = curr
    }
    if curSum > maxSum {
        return curSum
    }
    return maxSum
}