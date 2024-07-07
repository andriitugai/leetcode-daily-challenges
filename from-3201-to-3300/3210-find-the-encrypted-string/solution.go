func getEncryptedString(s string, k int) string {
    n := len(s)
    result := make([]byte, n)
    for i := 0; i < n; i ++ {
        result[i] = s[(i + k) % n]
    }
    return string(result)
}