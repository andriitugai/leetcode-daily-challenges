func winnerOfGame(colors string) bool {
    alice, bob := 0, 0
    left := 0
    for right := 0; right < len(colors); right ++ {
        if colors[left] != colors[right] {
            left = right
        }
        extra := right - left - 1
        if extra > 0 {
            if colors[right] == 'A' {
                alice += 1
            } else {
                bob += 1
            }
        }
    }
    return alice > bob
}