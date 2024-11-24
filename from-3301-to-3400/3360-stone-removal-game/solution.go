func canAliceWin(n int) bool {
    toRemove := 10
    round := 0
    for n >= toRemove {
        n -= toRemove
        round += 1
        toRemove -= 1
    }
    return round % 2 == 1
}