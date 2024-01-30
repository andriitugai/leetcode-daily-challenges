func minimumArrayLength(nums []int) int {
    minNum := 1000000001
    cntMin := 0
    for _, num := range nums {
        if num < minNum {
            minNum = num
            cntMin = 1
        } else if num == minNum {
            cntMin += 1
        }
    }

    for _, num := range nums {
        if num % minNum > 0 {
            return 1
        }
    }
    return (cntMin + 1) / 2
}