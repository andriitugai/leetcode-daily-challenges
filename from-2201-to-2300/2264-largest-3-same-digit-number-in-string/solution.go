func largestGoodInteger(num string) string {
    result := ""
    for i := 0; i < len(num) - 2; i ++ {
        if num[i] == num[i+1] && num[i] == num[i+2] && num[i:i+3] > result {
            result = num[i:i+3]
        }
    }
    return result
}