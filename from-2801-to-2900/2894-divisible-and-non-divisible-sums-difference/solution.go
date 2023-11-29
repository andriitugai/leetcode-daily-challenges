func differenceOfSums(n int, m int) int {
    result := 0
    for i := 1; i <= n; i ++ {
        if i % m != 0 {
            result += i
        } else {
            result -= i
        }
    }
    return result
}