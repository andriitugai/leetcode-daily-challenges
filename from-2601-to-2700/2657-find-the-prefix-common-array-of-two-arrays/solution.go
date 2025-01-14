func findThePrefixCommonArray(A []int, B []int) []int {
    common := 0
    n := len(A)
    result := make([]int, n)
    counts := make(map[int]int)
    for i := 0; i < n; i ++ {
        counts[A[i]] += 1
        if counts[A[i]] == 2 {
            common += 1
        }
        counts[B[i]] += 1
        if counts[B[i]] == 2 {
            common += 1
        }

        result[i] = common
    }
    return result
}