func returnToBoundaryCount(nums []int) int {
    result := 0
    curPos := 0
    for _, num := range nums {
        curPos += num
        if curPos == 0 {
            result += 1
        }
    }
    return result
}