func maximumElementAfterDecrementingAndRearranging(arr []int) int {
    sort.Ints(arr)
    result := 1
    for i := 1; i < len(arr); i ++ {
        if arr[i] > result {
            result += 1
        }
    }
    return result
}