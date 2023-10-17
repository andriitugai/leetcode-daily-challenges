func romanToInt(s string) int {
    romans := map[string]int {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
    result := 0
    var curr int
    prev := 0
    for i := len(s)-1; i >= 0; i-- {
        curr = romans[string(s[i])]
        if curr >= prev {
            result += curr
            prev = curr
        } else {
            result -= curr
        }
    }
    return result
}