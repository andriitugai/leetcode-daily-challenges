func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    result := 0
    for _, h := range hours {
        if h >= target {
            result += 1
        }
    }
    return result
}