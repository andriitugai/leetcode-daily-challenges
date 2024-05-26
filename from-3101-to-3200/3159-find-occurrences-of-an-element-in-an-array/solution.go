func occurrencesOfElement(nums []int, queries []int, x int) []int {
    idxs := make(map[int]int)
    for oc, i := 0, 0; i < len(nums); i ++ {
        if nums[i] == x {
            oc += 1
            idxs[oc] = i
        }
    }
    result := make([]int, len(queries))
    for i, q := range queries {
        if r, ok := idxs[q]; ok {
            result[i] = r
        } else {
            result[i] = -1
        }
    }
    return result
}