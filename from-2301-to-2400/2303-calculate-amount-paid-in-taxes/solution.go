func calculateTax(brackets [][]int, income int) float64 {
    var taxTotal float64 = 0.0
    prevBound := 0
    for i := 0; i < len(brackets) && income > 0; i++ {
        bound, ratio := brackets[i][0], brackets[i][1]
        taxed := bound - prevBound
        if taxed > income {
            taxed = income
        }
        income -= taxed
        prevBound = bound
        taxTotal += float64(ratio) * float64(taxed) / 100.0
    }
    return taxTotal
}