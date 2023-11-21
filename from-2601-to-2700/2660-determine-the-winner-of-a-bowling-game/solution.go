func isWinner(player1 []int, player2 []int) int {
    s1 := 0
    for i := 0; i < len(player1); i++ {
        s1 += player1[i]
        if i > 0 && player1[i-1] == 10 || i > 1 && player1[i - 2] == 10 {
            s1 += player1[i]
        }
    }
    s2 := 0
    for i := 0; i < len(player2); i++ {
        s2 += player2[i]
        if i > 0 && player2[i-1] == 10 || i > 1 && player2[i - 2] == 10 {
            s2 += player2[i]
        }
    }
    if s1 > s2 {
        return 1
    } else if s2 > s1 {
        return 2
    }
    return 0
}