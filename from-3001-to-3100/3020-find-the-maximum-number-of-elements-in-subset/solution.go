func maximumLength(nums []int) int {
    sort.Slice(nums, func(i, j int) bool{
        return nums[i] > nums[j]
    })

    counts := make(map[int]int)
    ones := 0
    for _, num := range nums {
        if num == 1 {
            ones += 1
        }
        if counts[num * num] >= 1 && counts[num] >= 1 && num != 1 {
            counts[num] = counts[num * num] + 2
        } else {
            counts[num] = 1
        }
    }
    
    maxVal := 0
    for _, val := range counts {
        if val >= maxVal {
            maxVal = val
        }
    }
    if ones % 2 == 0 {
        ones -= 1
    }
    if ones > maxVal {
        return ones
    }
    return maxVal
}