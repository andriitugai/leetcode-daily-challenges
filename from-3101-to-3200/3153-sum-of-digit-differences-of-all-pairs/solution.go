func sumDigitDifferences(nums []int) int64 {
    n := len(nums)
    figs := make([]string, n)
    pairs := (n -1) * n / 2

    for i := 0; i < n; i ++ {
        figs[i] = strconv.Itoa(nums[i])
    }
    figLen := len(figs[0])
    var result int64 = 0

    for idx := 0; idx < figLen; idx ++ {
        freq := make(map[byte]int)
        for _, fig := range figs {
            freq[fig[idx]] += 1
        }
        same := 0
        for _, v := range freq {
            if v > 1 {
                same += v * (v - 1) / 2
            }
        }
        result += int64(pairs - same)
    }
    return result
}