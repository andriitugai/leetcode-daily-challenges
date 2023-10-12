func decrypt(code []int, k int) []int {
    n := len(code)
    result := make([]int, n)
    if k == 0 {
        return result
    }
    start, end := 1, k + 1
    if k < 0 {
        start, end  =  n + k, n
    }
    s := 0
    for i := start; i < end; i++ {
        s += code[i]
    }
    for i := 0; i < n; i++ {
        result[i] = s
        s = s - code[(start + i) % n] + code[(end + i) % n]
    }
    return result
}