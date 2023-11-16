func countDigits(num int) int {
    result := 0
    num2 := num
    for num2 > 0 {
        if num % (num2 % 10) == 0 {
            result += 1
        }
        num2 /= 10
    }
    return result
}