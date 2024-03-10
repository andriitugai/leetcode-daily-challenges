func maximumHappinessSum(happiness []int, k int) int64 {
    sort.Slice(happiness, func(i, j int) bool{
        return happiness[j] < happiness[i]
    })
    var result int64 = 0
    turn := 0
    for i := 0; turn < k && i < len(happiness); i ++ {
        if happiness[i] > turn {
            result += int64(happiness[i] - turn)
        }
        turn += 1
    }
    return result
}