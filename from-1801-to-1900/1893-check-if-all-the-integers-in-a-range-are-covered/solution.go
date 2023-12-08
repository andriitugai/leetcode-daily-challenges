func isCovered(ranges [][]int, left int, right int) bool {
    covered := make(map[int]bool)
    for _, rng := range ranges {
        for i := rng[0]; i <= rng[1]; i ++ {
            covered[i] = true
        }
    }

    for i := left; i <= right; i ++ {
        if !covered[i] {
            return false
        }
    }
    return true
}