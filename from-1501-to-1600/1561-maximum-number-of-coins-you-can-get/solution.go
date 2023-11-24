func maxCoins(piles []int) int {
    sort.Ints(piles)
    result := 0
    n := len(piles)
    for i := n - 2; i >= n / 3; i -= 2 {
        result += piles[i]
    }
    return result
}