func numberOfWays(corridor string) int {
    x := 1
    y := 0
    z := 0
    mod := 1000000007
    for i := 0; i < len(corridor); i ++ {
        if corridor[i] == 'S' {
            x = (x + z) % mod
            z = y
            y = x
            x = 0
        } else {
            x = (x + z) % mod
        }
    }
    return z
}