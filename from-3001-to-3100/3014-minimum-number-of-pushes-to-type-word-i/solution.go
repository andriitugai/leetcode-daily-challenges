func minimumPushes(word string) int {
    cnt := make(map[rune]int)
    for _, r := range word {
        cnt[r] += 1
    }
    vals := make([]int, len(cnt))
    for _, v := range cnt {
        vals = append(vals, v)
    }
    sort.Slice(vals, func(i, j int) bool {
        return vals[i] > vals[j]
    })
    result := 0
    for i := 0; i < len(vals); i ++ {
        result += (i / 8 + 1) * vals[i]
    }
    return result
}