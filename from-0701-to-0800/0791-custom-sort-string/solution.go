func customSortString(order string, s string) string {
    orderMap := make(map[byte]int)
    for i := 0; i < len(order); i ++ {
        orderMap[order[i]] = i + 1
    }
    bytes := []byte(s)
    sort.Slice(bytes, func(i, j int) bool {
        return orderMap[bytes[i]] < orderMap[bytes[j]]
    })
    return string(bytes)
}