func fillCups(amount []int) int {
    result := amount[0]
    total := (amount[0] + amount[1] + amount[2] + 1) / 2
    if amount[1] > result {
        result = amount[1]
    }
    if amount[2] > result {
        result = amount[2]
    }
    if total > result {
        result = total
    }
    return result   
}