func minimumOneBitOperations(n int) int {
    result := 0
    for n > 0 {
        result = - result - (n ^ (n - 1))
        n &= n - 1
    }
    if result < 0 {
        return -result
    }
    return result
}