func judgeSquareSum(c int) bool {
    left := 0
    right := int(math.Floor(math.Sqrt(float64(c))))
    var res int

    for left <= right {
        res = left * left + right * right
        if res == c {
            return true
        } else if res < c {
            left += 1
        } else {
            right -= 1
        }
    }
    return false
}