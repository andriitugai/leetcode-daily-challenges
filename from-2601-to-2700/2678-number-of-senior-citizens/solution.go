func countSeniors(details []string) int {
    result := 0
    for _, d := range details {
        if (d[11] == '6' && d[12] > '0') || d[11] > '6' {
            result += 1
        }
    }
    return result
}