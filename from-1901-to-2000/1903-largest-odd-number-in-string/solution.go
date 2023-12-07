func largestOddNumber(num string) string {
    r := len(num) - 1
    for r >= 0 {
        if int(num[r] - '0') % 2 == 1 {
            break
        }
        r -= 1
    }
    return num[:r+1]
}