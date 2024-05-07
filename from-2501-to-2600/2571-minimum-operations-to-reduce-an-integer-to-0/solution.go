func minOperations(n int) int {
    result := 0
    for n != 0 {
        n = n - closestPow2(n)
        if n < 0 {
            n = -n
        }
        result += 1
    }
    return result
}

func closestPow2(num int) int {
    upper := 1
    for upper < num {
        upper <<= 1
    }
    lower := upper >> 1
    if upper - num < num - lower {
        return upper
    }
    return lower
}