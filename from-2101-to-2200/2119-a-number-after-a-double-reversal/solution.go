func isSameAfterReversals(num int) bool {
    reverse := func(n int) int {
        result := 0
        for n > 0 {
            result = result * 10 + n % 10
            n /= 10
        }
        return result
    }
    return num == reverse(reverse(num))
}