func monkeyMove(n int) int {
    mod := 1000000007
    p := 1
    exp := 2
    for n > 0 {
        if n & 1 != 0 {
            p *= exp
            p %= mod
        }
        exp = (exp * exp) % mod
        n >>= 1
    }
    return (p - 2 + mod) % mod
}