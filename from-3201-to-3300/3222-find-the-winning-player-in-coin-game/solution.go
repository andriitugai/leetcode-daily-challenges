func losingPlayer(x int, y int) string {
    numOfSets := y / 4
    if x < numOfSets {
        numOfSets = x
    }
    if numOfSets % 2 == 1 {
        return "Alice"
    }
    return "Bob"
}