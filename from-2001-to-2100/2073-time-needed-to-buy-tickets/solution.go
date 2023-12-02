func timeRequiredToBuy(tickets []int, k int) int {
    result := 0
    i := 0
    for true {
        if tickets[i] > 0 {
            tickets[i] -= 1
            result += 1
            if i == k && tickets[i] == 0 {
                break
            }
        }
        i += 1
        if i == len(tickets) {
            i = 0
        }
    }
    return result
}