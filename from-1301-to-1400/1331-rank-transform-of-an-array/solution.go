type item struct {
    val, idx int
}

func arrayRankTransform(arr []int) []int {
    items := make([]item, len(arr))
    for i := 0; i < len(arr); i++ {
        items[i] = item{arr[i], i}
    }

    sort.Slice(items, func(i, j int) bool {
        return items[i].val < items[j].val
    })

    result := make([]int, len(arr))
    seen := make(map[int]bool)
    rank := 0
    for _, el := range items {
        if !seen[el.val] {
            seen[el.val] = true
            rank ++
        }
        result[el.idx] = rank
     }
     return result
}