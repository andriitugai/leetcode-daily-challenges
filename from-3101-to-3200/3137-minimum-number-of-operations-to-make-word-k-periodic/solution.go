func minimumOperationsToMakeKPeriodic(word string, k int) int {
    countSub := make(map[string]int)
    maxSub := 0
    for i := 0; i < len(word); i += k {
        sub := word[i:i+k]
        countSub[sub] += 1
        if countSub[sub] > maxSub {
            maxSub = countSub[sub]
        }
    }
    return len(word) / k - maxSub
}