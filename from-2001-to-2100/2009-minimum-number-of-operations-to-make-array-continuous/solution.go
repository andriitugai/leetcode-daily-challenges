func minOperations(nums []int) int {
    n := len(nums)
    result := n
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }

    set := make(map[int]bool)
    for _, num := range nums {
        set[num] = true
    }

    unique := []int{}
    for num, _ := range set {
        unique = append(unique, num)
    }
    sort.Ints(unique)

    j := 0
    m := len(unique)

    for i := 0; i < m; i++ {
        for j < m && unique[j] < unique[i] + n {
            j += 1
        }
        result = min(result, n - j + i)
    }
    return result
}