func buyChoco(prices []int, money int) int {
    c1, c2 := 150, 150
    for _, p := range prices {
        if c1 > p {
            c2 = c1
            c1 = p
        } else if c2 > p {
            c2 = p
        }
    }
    
    if c1 + c2 > money {
        return money
    }
    return money - c1 - c2
}