func isThree(n int) bool {
    // The Newton's algorithm to determine if n is a perfect square
    // x1, x2 := 1, n
    // for x2 > x1 {
    //     x2 = (x1 + x2) / 2
    //     x1 = n / x2
    // }
    // return x1 == x2 && n % x1 == 0
    if n < 4 {
        return false
    }

    for k := 2; k * k <= n; k ++ {
        if n % k == 0 && k * k != n {
            return false
        }
        if k * k == n {
            return true
        }
    }
    return false
}