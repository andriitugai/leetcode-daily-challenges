func findRelativeRanks(score []int) []string {
    idx := make(map[int]int)
    for i := 0; i < len(score); i ++ {
        idx[score[i]] = i
    }
    
    sort.Slice(score, func(i, j int) bool {
        return score[i] > score[j]
    })

    result := make([]string, len(score))
    for i := 0; i < len(score); i ++ {
        if i == 0 {
            result[idx[score[i]]] = "Gold Medal"
        } else if i == 1 {
            result[idx[score[i]]] = "Silver Medal"
        } else if i == 2 {
            result[idx[score[i]]] = "Bronze Medal"
        } else {
            result[idx[score[i]]] = strconv.Itoa(i + 1)
        }
    }
    return result
}