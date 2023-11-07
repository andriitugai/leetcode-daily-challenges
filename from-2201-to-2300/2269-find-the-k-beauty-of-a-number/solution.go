unc divisorSubstrings(num int, k int) int {
    str := strconv.Itoa(num)

    result := 0

    for i := range str {
        n, _ := strconv.Atoi(str[i:i+k])
        if n != 0 && num % n == 0 {
            result += 1
        }
        if i + k == len(str) {
            break
        }
    }
    return result
}