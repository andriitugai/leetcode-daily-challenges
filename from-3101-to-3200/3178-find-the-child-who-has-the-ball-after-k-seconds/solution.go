func numberOfChild(n int, k int) int {
    dir := 1
    curPos := 0
    for i := 0; i < k; i ++ {
        curPos += dir
        if curPos == 0 || curPos == n - 1 {
            dir *= -1
        }
    }
    return curPos
}