func frequencySort(s string) string {
    freq := make(map[byte]int)
    for i := 0; i < len(s); i ++ {
        freq[s[i]] += 1
    }
    items := [][]int{}
    for c, f := range freq {
        items = append(items, []int{int(c), f})
    }
    sort.Slice(items, func(i, j int) bool {
        return items[i][1] > items[j][1]
    })
    result := []byte{}
    for _, item := range items {
        for i := 0; i < item[1]; i++ {
            result = append(result, byte(item[0]))
        }
    }
    return string(result)
}