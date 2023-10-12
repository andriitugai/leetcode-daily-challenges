func canFormArray(arr []int, pieces [][]int) bool {
    m := make(map[int][]int)
    for _, p := range pieces {
        m[p[0]] = p
    }

    result := []int{}
    for _, v := range arr {
        if p, ok := m[v]; ok {
            result = append(result, p...)
        }
    }

    for i, v := range arr {
        if i >= len(result)|| v != result[i] {
            return false
        }
    }
    return true
}