func isBalanced(num string) bool {
    evenSum, oddSum := 0, 0
    var digit int
    for i := 0; i < len(num); i ++ {
        digit = int(num[i] - '0')
        if i % 2 == 0 {
            evenSum += digit
        } else {
            oddSum += digit
        }
    }
    return evenSum == oddSum
}