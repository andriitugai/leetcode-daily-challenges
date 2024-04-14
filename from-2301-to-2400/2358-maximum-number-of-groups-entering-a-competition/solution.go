func maximumGroups(grades []int) int {
    sort.Ints(grades)
    nGroups := 0
    n := len(grades)
    gSize := 1
    curSum, prevSum := 0, 0
    idx := 0
    for n > 0 {
        prevSum = curSum
        curSum = 0
        for i := idx; i < len(grades) && i < idx + gSize; i ++ {
            curSum += grades[i]
        }
        n -= gSize
        idx += gSize
        gSize += 1
        nGroups += 1
    }
    if curSum <= prevSum || idx > len(grades){
        nGroups -= 1
    }
    return nGroups
}