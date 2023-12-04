func getLucky(s string, k int) int {
    sum := 0
    for i := 0; i < len(s); i ++ {
        n := int(s[i] - 'a' + 1)
        for n > 0 {
            sum += n % 10
            n /= 10
        }
    }

    for j := 1; j < k; j++ {
        newSum := 0
        for sum > 0 {
            newSum += sum % 10
            sum /= 10
        }
        sum = newSum
    }

    return sum
}