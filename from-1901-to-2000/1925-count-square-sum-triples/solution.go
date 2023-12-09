func countTriples(n int) int {
    squares := make(map[int]bool)
    for i := 0; i <= n; i ++ {
        squares[i * i] = true
    }

    result := 0

    for a := 3; a <= n - 2; a ++ {
        for b := a + 1; b <= n - 1; b ++ {
            if squares[a * a + b * b] {
                result += 2
            }
        }
    }
    return result
}