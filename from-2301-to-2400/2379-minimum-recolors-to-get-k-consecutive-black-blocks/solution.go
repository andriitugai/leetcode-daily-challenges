func minimumRecolors(blocks string, k int) int {
    cWhite := 0
    p0, p1 := 0, 0

    for p1 < k {
        if blocks[p1] == 'W' {
            cWhite += 1
        }
        p1 += 1
    }
    minWhite := cWhite

    for p1 < len(blocks) {
        if blocks[p1] == 'W' {
            cWhite += 1
        }
        p1 += 1
        if blocks[p0] == 'W' {
            cWhite -= 1
        }
        p0 += 1
        if cWhite < minWhite {
            minWhite = cWhite
        }
    }
    return minWhite
}