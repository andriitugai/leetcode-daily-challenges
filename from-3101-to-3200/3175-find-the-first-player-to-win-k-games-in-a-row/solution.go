func findWinningPlayer(skills []int, k int) int {
    n := len(skills)
    if k >= n {
        maxVal := skills[0]
        maxIdx := 0
        for i := 1; i < n; i ++ {
            if skills[i] > maxVal {
                maxVal = skills[i]
                maxIdx = i
            }
        }
        return maxIdx
    }
    winner := skills[0]
    winIdx := 0
    inRow := 1
    if skills[0] < skills[1] {
        winner = skills[1]
        winIdx = 1
    }
    for i := 2; i < n && inRow < k; i ++ {
        if winner > skills[i] {
            inRow += 1
        } else {
            winner = skills[i]
            winIdx = i
            inRow = 1
        }
    }
    return winIdx
}