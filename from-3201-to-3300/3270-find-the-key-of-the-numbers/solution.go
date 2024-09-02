func generateKey(num1 int, num2 int, num3 int) int {
    result := 0
    mod := 1000
    for i := 0; i < 4; i ++ {
        result = result * 10 + min(num1 / mod, num2 / mod, num3 / mod)
        num1 %= mod
        num2 %= mod
        num3 %= mod
        mod /= 10
    }
    return result
}

func min(a, b, c int) int {
    if a < b {
        if c < a {
            return c
        }
        return a
    } else {
        if c < b {
            return c
        }
    }
    return b
}