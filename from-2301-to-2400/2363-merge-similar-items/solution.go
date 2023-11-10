func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
    sort.Slice(items1, func(i, j int) bool {
        return items1[i][0] < items1[j][0]
    })
    sort.Slice(items2, func(i, j int) bool {
        return items2[i][0] < items2[j][0]
    })
    p1, p2 := 0, 0
    result := [][]int{}
    
    for p1 < len(items1) && p2 < len(items2) {
        if items1[p1][0] == items2[p2][0] {
            result = append(result, []int{items1[p1][0], items1[p1][1] + items2[p2][1]})
            p1 += 1
            p2 += 1
        } else if items1[p1][0] < items2[p2][0] {
            result = append(result, items1[p1])
            p1 += 1
        } else {
            result = append(result, items2[p2])
            p2 += 1
        }
    }

    for p1 < len(items1) {
        result = append(result, items1[p1])
        p1 += 1
    }

    for p2 < len(items2) {
        result = append(result, items2[p2])
        p2 += 1
    }

    return result
}