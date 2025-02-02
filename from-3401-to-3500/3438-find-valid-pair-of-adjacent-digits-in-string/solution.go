func findValidPair(s string) string {
    freq := make(map[int]int)
    n := len(s)
    nums := make([]int, n)
    for i, c := range s {
        nums[i] = int(c - '0')
        freq[nums[i]] += 1
    }
    for i := 1; i < n; i ++ {
        if nums[i - 1] != nums[i] && freq[nums[i - 1]] == nums[i - 1] && freq[nums[i]] == nums[i] {
            return s[i - 1: i + 1]
        }
    }
    return ""
}