func findMinimumOperations(s1 string, s2 string, s3 string) int {
    if s1[0] != s2[0] || s1[0] != s3[0] || s2[0] != s3[0] {
        return -1
    }
    n1, n2, n3 := len(s1), len(s2), len(s3)
    i := 1
    for i < n1 && i < n2 && i < n3 {
        if s1[i] != s2[i] || s2[i] != s3[i] || s1[i] != s3[i] {
            break
        }
        i += 1
    }
    return n1 - i + n2 - i + n3 - i
}