func countVowelPermutation(n int) int {
    mod := 1000000007
    a, e, i, o, u := 0, 1, 2, 3, 4
    prev := []int{1, 1, 1, 1, 1}

    for idx := 0; idx < n -1; idx++ {
        curr := make([]int, 5)
        curr[a] = (prev[e] + prev[i] + prev[u]) % mod
        curr[e] = (prev[i] + prev[a]) % mod
        curr[i] = (prev[e] + prev[o]) % mod
        curr[o] = prev[i]
        curr[u] = (prev[i] + prev[o]) % mod

        prev = curr
    }

    result := 0
    for idx := 0; idx < 5; idx ++ {
        result = (result + prev[idx]) % mod
    }
    return result
}