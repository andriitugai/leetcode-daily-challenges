func lastRemaining(n int) int {
    var lastR func(num int) int
    lastR = func(num int) int {
        if num == 1 {
            return 1
        }
        return 2 * (num / 2 + 1 - lastR(num / 2))
    }
    return lastR(n)
}