func getSum(a int, b int) int {
    var sum func(a, b int) int
    sum = func(a, b int) int {
        if b == 0 {
            return a
        }
        return sum(a ^ b, (a & b) << 1)
    }
    return sum(a, b)
}