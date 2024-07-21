func minChanges(n int, k int) int {
    if k > n {
        return -1
    } else if k == n {
        return 0
    }
    result := 0
    for n > 0 {
        if k % 2 == 0 {
            if n % 2 == 1 {
                result += 1
            }
        } else {
            if n % 2 == 0 {
                return -1
            }
        }
        k /= 2
        n /= 2
    }
    return result
}