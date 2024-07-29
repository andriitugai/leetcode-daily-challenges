func numTeams(rating []int) int {
    n := len(rating)
    result := 0

    for j := 0; j < n; j ++ {
        rLess, rGreater, lLess, lGreater := 0, 0, 0, 0
        for i := 0; i < j; i ++ {
            if rating[i] < rating[j] {
                lLess += 1
            } else if rating[i] > rating[j] {
                lGreater += 1
            }
        }
        for k := j + 1; k < n; k ++ {
            if rating[k] < rating[j] {
                rLess += 1
            } else if rating[k] > rating[j] {
                rGreater += 1
            }
        }
        result += lLess * rGreater + lGreater * rLess
    }
    return result
}