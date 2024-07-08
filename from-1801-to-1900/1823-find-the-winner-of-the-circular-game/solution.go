func findTheWinner(n int, k int) int {
    players := make([]int, n)
    var player int
    for i := 0; i < n; i ++ {
        players[i] = i + 1
    }
    count := 0
    for len(players) > 0 {
        player = players[0]
        players = players[1:]
        count += 1
        if count == k {
            count = 0
        } else {
            players = append(players, player)
        }
    }
    return player
}