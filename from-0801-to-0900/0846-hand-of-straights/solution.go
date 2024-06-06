func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand) % groupSize != 0 {
        return false
    }
    count := make(map[int]int)
    for _, card := range hand {
        count[card] += 1
    }
    cards := []int{}
    for card, _ := range count {
        cards = append(cards, card)
    }
    sort.Ints(cards)

    for _, card := range cards {
        if count[card] > 0 {
            startCount := count[card]
            for i := card; i < card + groupSize; i ++ {
                if count[i] < startCount {
                    return false
                }
                count[i] -= startCount
            }
        }
    }
    return true
}