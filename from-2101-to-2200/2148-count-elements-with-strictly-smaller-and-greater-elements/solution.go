func countElements(nums []int) int {
    minEl, maxEl := 1000000, -1000000
    minCnt, maxCnt := 0, 0

    for _, num := range nums {
        if num < minEl {
            minEl = num
            minCnt = 1
        } else if num == minEl {
            minCnt += 1
        } 
        if num > maxEl {
            maxEl = num
            maxCnt = 1
        } else if num == maxEl {
            maxCnt += 1
        }
    }
    result := len(nums) - minCnt - maxCnt
    if result < 0 {
        result = 0
    }
    return result
}