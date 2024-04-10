func deckRevealedIncreasing(deck []int) []int {
    sort.Ints(deck)

    q := make([]int, len(deck))
    for i := 0; i < len(deck); i ++ {
        q[i] = i
    }

    var idx int
    result := make([]int, len(deck))

    for _, card := range deck {
        idx = q[0]
        result[idx] = card
        q = q[1:]

        if len(q) > 0 {
            idx = q[0]
            q = q[1:]
            q = append(q, idx)
        }
    }
    return result
}