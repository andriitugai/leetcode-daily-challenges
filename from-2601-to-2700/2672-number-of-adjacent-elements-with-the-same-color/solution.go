func colorTheArray(n int, queries [][]int) []int {
    arr := make([]int, n)
    ans := 0
    result := make([]int, len(queries))
    var idx, clr int

    for i, q := range queries {
        idx, clr = q[0], q[1]
        if idx > 0 {
            if arr[idx-1] == clr && arr[idx-1] != arr[idx] {
                ans += 1
            } else if arr[idx-1] > 0 && arr[idx-1] != clr && arr[idx-1] == arr[idx] && ans > 0 {
                ans -= 1
            }
        }
        if idx < n - 1 {
            if arr[idx+1] == clr && arr[idx+1] != arr[idx] {
                ans += 1
            } else if arr[idx+1] > 0 && arr[idx+1] == arr[idx] && arr[idx+1] != clr && ans > 0 {
                ans -= 1
            }
        }
        arr[idx] = clr
        result[i] = ans
    }
    return result
}