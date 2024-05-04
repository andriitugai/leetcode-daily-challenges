func findNthDigit(n int) int {
    number, count := 1, 9
    p := 1
    for count * p < n {
        n -= count * p
        count *= 10
        number *= 10
        p += 1
    }
    number += (n - 1) / p
    digit := (n - 1) % p
    
    return int(strconv.Itoa(number)[digit] - '0')
}