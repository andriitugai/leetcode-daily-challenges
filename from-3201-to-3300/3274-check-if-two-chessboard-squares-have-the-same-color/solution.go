func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
    // 0, 0 or 1, 1  - black (1)
    // 0, 1 or 1, 0  - white (0)
    getColor := func(c string) int {
        if (c[0] - 'a') % 2 == (c[1] - '0') % 2 {
            return 1
        }
        return 0
    }
    if getColor(coordinate1) == getColor(coordinate2) {
            return true
    }
    return false
}