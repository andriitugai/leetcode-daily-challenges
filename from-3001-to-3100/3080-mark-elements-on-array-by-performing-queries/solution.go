type ElemInfo struct {
    Idx int
    Val int64
    Flag bool
}

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
    elInfo := make([]*ElemInfo, len(nums))
    elMap := make(map[int]*ElemInfo)
    var total int64 = 0
    totals := make([]int64, len(queries))

    for idx, num := range nums {
        info := &ElemInfo{Idx: idx, Val: int64(num), Flag: false}
        elInfo[idx] = info
        elMap[idx] = info
        total += int64(num)
    }

    for i := 0; i < len(totals); i ++ {
        totals[i] = total
    }

    sort.Slice(elInfo, func(i, j int) bool {
        if elInfo[i].Val == elInfo[j].Val {
            return elInfo[i].Idx < elInfo[j].Idx
        }
        return elInfo[i].Val < elInfo[j].Val
    })

    minIdx, n := 0, len(nums)
    for qix, q := range queries {
        i, k := q[0], q[1]
        if elMap[i].Flag == false {
            elMap[i].Flag = true
            total -= elMap[i].Val
        }
        for k > 0 {
            for minIdx < n && elInfo[minIdx].Flag == true {
                minIdx += 1
            }
            if minIdx < n {
                elInfo[minIdx].Flag = true
                total -= elInfo[minIdx].Val
                minIdx += 1
            }
            k -= 1
        }
        totals[qix] = total
    }
    return totals
}