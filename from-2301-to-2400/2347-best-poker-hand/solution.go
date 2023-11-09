func bestHand(ranks []int, suits []byte) string {
    // Flush?
    isFlush := true
    for i := 1; i < len(suits); i++ {
        if suits[i] != suits[0] {
            isFlush = false
            break
        }
    }
    if isFlush {
        return "Flush"
    }
    cnt := make(map[int]int)
    isThree := false
    isPair := false
    for _, r := range ranks {
        cnt[r] += 1
    }
    for _, c := range cnt {
        if c >= 3 {
            isThree = true
        }
        if c == 2 {
            isPair = true
        }
    }
    if isThree {
        return "Three of a Kind"
    }
    if isPair {
        return "Pair"
    }
    return "High Card"
}