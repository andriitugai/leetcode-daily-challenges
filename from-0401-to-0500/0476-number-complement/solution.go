func findComplement(num int) int {
    result := 0
    mask := 1
    for num > 0 {
        if num % 2 == 0 {
            result += mask
        }
        num /= 2
        mask <<= 1
    }
    return result
}