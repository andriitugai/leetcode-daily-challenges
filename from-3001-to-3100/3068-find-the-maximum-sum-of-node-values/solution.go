func maximumValueSum(nums []int, k int, edges [][]int) int64 {
    var total int64 = 0
    for _, num := range nums {
        total += int64(num)
    }
    var totalDiff int64 = 0
    positiveCount := 0
    minAbsDiff := math.MaxInt32
    var diff int

    for _, num := range nums {
        diff = (num ^ k) - num

        if diff > 0 {
            totalDiff += int64(diff)
            positiveCount += 1
        }
        minAbsDiff = min(minAbsDiff, abs(diff))
    }

    if positiveCount % 2 == 1 {
        totalDiff -= int64(minAbsDiff)
    }

    return total + totalDiff
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}