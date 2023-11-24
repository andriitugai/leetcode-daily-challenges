func isFascinating(n int) bool {
    if n > 333 {
        return false
    }
    digits := make(map[byte]bool)
    r := strconv.Itoa(n) + strconv.Itoa(2 * n) + strconv.Itoa(3 * n)
    for i := 0; i < len(r); i ++ {
        if r[i] != '0' {
            digits[r[i]] = true
        }
    }
    return len(digits) == 9
}