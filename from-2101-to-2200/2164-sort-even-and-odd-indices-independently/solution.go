func sortEvenOdd(nums []int) []int {
    odds, evens := []int{}, []int{}
    for i := 0; i < len(nums); i ++ {
        if i % 2 == 1 {
            odds = append(odds, nums[i])
        } else {
            evens = append(evens, nums[i])
        }
    }

    sort.Slice(odds, func(i, j int) bool {
        return odds[i] > odds[j]
    })
    sort.Slice(evens, func(i, j int) bool {
        return evens[i] < evens[j]
    })
    result := []int{}
    pe, po := 0,0
    for po < len(odds) && pe < len(evens) {
        result = append(result, evens[pe])
        result = append(result, odds[po])
        po += 1
        pe += 1
    }
    if po < len(odds) {
        result = append(result, odds[po])
    }
    if pe < len(evens) {
        result = append(result, evens[pe])
    }
    return result
}