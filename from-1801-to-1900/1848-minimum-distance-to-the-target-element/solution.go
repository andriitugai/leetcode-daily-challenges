func getMinDistance(nums []int, target int, start int) int {
    result := 1000000000
    for i := 0; i < len(nums); i ++ {
        if nums[i] == target {
            dist := i - start
            if dist < 0 {
                dist = -dist
            }
            if dist < result {
                result = dist
            }
        }
    }
    return result
}