func minimumDeletions(word string, k int) int {
    counts := make(map[byte]int)
    for i := 0; i < len(word); i ++ {
        counts[word[i]] += 1
    }
    freqs := []int{}
    for _, f := range counts {
        freqs = append(freqs, f)
    }
    sort.Ints(freqs)
    n := len(freqs)
    deleted, ans := 0, len(word)

    for i := 0; i < n; i ++ {
        res := deleted
        minFreq := freqs[i]
        for j := n - 1; j > i; j -- {
            if freqs[j] - minFreq <= k {
                break
            }
            res += freqs[j] - minFreq - k
        }
        if res < ans {
            ans = res
        }
        deleted += freqs[i]
    }
    return ans
}