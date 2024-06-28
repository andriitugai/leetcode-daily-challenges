func maximumImportance(n int, roads [][]int) int64 {
    edgeCount := make([]int, n)
    for _, r := range roads {
        edgeCount[r[0]] += 1
        edgeCount[r[1]] += 1
    }
    sort.Ints(edgeCount)
    var result int64 = 0
    label := 1
    for _, count := range edgeCount {
        result += int64(label * count)
        label += 1
    }
    return result
}