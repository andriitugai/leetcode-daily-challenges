func splitNum(num int) int {
    numbers := make([]int, 10)
    for num > 0 {
        numbers[num % 10]++
        num /= 10
    }
    n1, n2 := 0, 0
    odd := true
    for i := 0; i < 10; i++ {
        for numbers[i] > 0 {
            if odd {
                n1 = n1 * 10 + i
            } else {
                n2 = n2 * 10 + i
            }
            numbers[i]--
            odd = !odd
        }
    }
    return n1 + n2    
}