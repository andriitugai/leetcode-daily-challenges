func findWinners(matches [][]int) [][]int {
    players := make(map[int]bool)
    loses := make(map[int]int)
    for _, match := range matches {
        players[match[0]] = true
        players[match[1]] = true
        loses[match[1]] += 1
    }
    lose_zero := []int{}
    lose_one := []int{}
    for p, _ := range players {
        if loses[p] == 0 {
            lose_zero = append(lose_zero, p)
        } else if loses[p] == 1 {
            lose_one = append(lose_one, p)
        }
    }
    sort.Ints(lose_zero)
    sort.Ints(lose_one)
    return [][]int{lose_zero, lose_one}
}