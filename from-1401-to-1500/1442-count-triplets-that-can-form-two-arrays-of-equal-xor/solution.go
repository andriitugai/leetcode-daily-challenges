func countTriplets(arr []int) int {
    n := len(arr)
    result := 0
    for i := 0; i < n - 1; i ++ {
        a := 0
        for j := i + 1; j < n; j ++ {
            a ^= arr[j - 1]
            b := 0
            for k := j; k < n; k ++ {
                b ^= arr[k]
                if a == b {
                    result += 1
                }
            }
        }
    }
    return result
}