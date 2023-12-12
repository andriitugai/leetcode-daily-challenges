func countTestedDevices(bp []int) int {
    result := 0
    for i := 0; i < len(bp); i ++ {
        if bp[i] > 0 {
            result += 1
            for j := i + 1; j < len(bp); j ++ {
                if bp[j] > 0 {
                    bp[j] -= 1
                }
            }
        }
    }
    return result
}