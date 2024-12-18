func finalPrices(prices []int) []int {
    result := make([]int, len(prices))

    for i := 0; i < len(prices); i++ {
        price := prices[i]
        for j := i+1; j < len(prices); j++ {
            if prices[j] <= price {
                price -= prices[j]
                break
            }
        }
        result[i] = price
    }
    return result
}