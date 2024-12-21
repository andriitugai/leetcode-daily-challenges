func isPerfectSquare(num int) bool {
    if num == 1 {
        return true
    }
    left, right := 1, int(num / 2)
    for left <= right {
        mid := left + (right - left) / 2
        m2 := mid * mid
        if m2 == num {
            return true
        } else if m2 > num {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return false
}