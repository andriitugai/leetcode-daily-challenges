func getWinner(arr []int, k int) int {
    winCount := make(map[int]int)
    curWin := arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] > curWin {
            curWin = arr[i]
        }
        winCount[curWin] += 1
        if winCount[curWin] >= k {
            return curWin
        }
    }
    return curWin
}