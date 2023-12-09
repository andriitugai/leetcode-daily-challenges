func maxProductDifference(nums []int) int {
    max1, max2 := 0, 0
    min1, min2 := 100000, 100000

    for i := 0; i < len(nums); i ++ {
        if nums[i] > max1 {
            max2 = max1
            max1 = nums[i]
        } else if nums[i] > max2 {
            max2 = nums[i]
        }
        if nums[i] < min1 {
            min2 = min1
            min1 = nums[i]
        } else if nums[i] < min2 {
            min2 = nums[i]
        }
    }
    
    return max1 * max2 - min1 * min2
}