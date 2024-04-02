func sumOfTheDigitsOfHarshadNumber(x int) int {
    sd := 0
    num := x
    for x > 0 {
        sd += x % 10
        x /= 10
    }
    if num % sd == 0 {
        return sd
    }
    return -1
}