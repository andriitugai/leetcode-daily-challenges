func sumOfNumberAndReverse(num int) bool {
    reverse := func(num int) int {
        result := 0
        for num > 0 {
            result = result * 10 + num % 10
            num /= 10
        }
        return result
    }

    for i := num; i >= num / 2; i -- {
        if i + reverse(i) == num {
            return true
        }
    }
    return false
}