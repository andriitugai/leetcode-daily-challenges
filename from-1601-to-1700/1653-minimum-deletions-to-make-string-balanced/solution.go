func minimumDeletions(s string) int {
    n := len(s)
    aRight := make([]int, n)
    aTot := 0
    for i := n - 2; i >= 0; i -- {
        if s[i + 1] == 'a' {
            aTot += 1
        }
        aRight[i] = aTot
    }
    minDel := 1000000
    bTot := 0
    var curDel int
    for i := 0; i < n; i ++ {
        curDel = bTot + aRight[i]
        if curDel < minDel {
            minDel = curDel
        }
        if s[i] == 'b' {
            bTot += 1
        }
    }
    return minDel
}