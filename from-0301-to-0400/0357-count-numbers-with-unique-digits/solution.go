func countNumbersWithUniqueDigits(n int) int {
    if n == 0 {
        return 1
    }
    result := 10
    uniqueDigits := 9
    availNumbers := 9

    for n > 1 && availNumbers > 0 {
        uniqueDigits = uniqueDigits * availNumbers
        result += uniqueDigits
        availNumbers -= 1
        n -= 1
    }

    return result
}