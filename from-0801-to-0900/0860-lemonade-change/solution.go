func lemonadeChange(bills []int) bool {
    bank := map[int]int{5: 0, 10: 0, 20: 0}
    for _, bill := range bills {
        bank[bill] += 1
        change := bill - 5
        for change > 0 {
            for change >= 20 && bank[20] > 0 {
                bank[20] -= 1
                change -= 20
            }
            for change >= 10 && bank[10] > 0 {
                bank[10] -= 1
                change -= 10
            }
            for change >= 5 && bank[5] > 0 {
                bank[5] -= 1
                change -= 5
            }

            if change > 0 {
                return false
            }
        }
    }
    return true
}