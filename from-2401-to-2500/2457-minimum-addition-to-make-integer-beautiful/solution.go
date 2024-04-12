func makeIntegerBeautiful(n int64, target int) int64 {
    digits := []int{}
    var d, dSum int = 0, 0
    num := n
    for n > 0 {
        d = int(n % 10)
        digits = append(digits, d)
        dSum += d
        n /= 10
    }

    for i := 0; i < len(digits); i ++ {
        if dSum <= target {
            break
        }
        dSum -= digits[i]
        digits[i] = 0
        if i + 1 < len(digits) {
            digits[i + 1] += 1
        } else {
            digits = append(digits, 1)
        }
        dSum += 1
    }
    var newNum int64 = 0
    for i := len(digits) - 1; i >= 0; i -- {
        newNum = newNum * 10 + int64(digits[i])
    }
    return newNum - num
}