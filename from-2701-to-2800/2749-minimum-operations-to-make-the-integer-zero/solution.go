func makeTheIntegerZero(num1 int, num2 int) int {
    result := 0
    for countSetBits(num1) > result {
        result += 1
        num1 -= num2
        if num1 < result && num2 > 0 {
            return -1
        }
    }
    return result
}

func countSetBits(n int) int {
    count := 0
    for n > 0 {
        n &= n - 1
        count += 1
    }
    return count
}