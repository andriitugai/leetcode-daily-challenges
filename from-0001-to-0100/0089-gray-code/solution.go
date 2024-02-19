func grayCode(n int) []int {
    result := make([]int, 1 << n)

    for i := 0; i < 1 << n; i ++ {
        result[i] = i ^ (i >> 1)
    }
    return result
}