func relativeSortArray(arr1 []int, arr2 []int) []int {
    count := make(map[int]int)
    for _, num := range arr1 {
        count[num] += 1
    }
    result := make([]int, len(arr1))
    pr := 0
    for i := 0; i < len(arr2); i ++ {
        for j := 0; j < count[arr2[i]]; j ++ {
            result[pr] = arr2[i]
            pr += 1
        }
        delete(count, arr2[i])
    }
    keys := []int{}
    for k, _ := range count {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    for i := 0; i < len(keys); i ++ {
        for j := 0; j < count[keys[i]]; j ++ {
            result[pr] = keys[i]
            pr += 1
        }
    }
    return result
}