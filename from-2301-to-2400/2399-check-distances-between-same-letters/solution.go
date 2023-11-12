func checkDistances(s string, distance []int) bool {
    positions := make(map[int][]int)
    for i := 0; i < len(s); i ++ {
        idx := int(s[i] - 'a')
        positions[idx] = append(positions[idx], i)
    }

    for lidx, px := range positions {
        if px[1] - px[0] - 1 != distance[lidx] {
            return false
        }
    }
    return true
}