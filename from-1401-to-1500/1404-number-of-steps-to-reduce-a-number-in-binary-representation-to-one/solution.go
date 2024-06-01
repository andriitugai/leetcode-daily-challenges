func numSteps(s string) int {
    result := 0
    carry := 0
    var digit int
    for i := len(s) - 1; i >= 1; i -- {
        d, _ := strconv.Atoi(string(s[i]))
        digit = (d + carry) % 2
        if digit == 0 {
            result += 1
        } else {
            result += 2
            carry = 1
        }
    }
    return result + carry
}