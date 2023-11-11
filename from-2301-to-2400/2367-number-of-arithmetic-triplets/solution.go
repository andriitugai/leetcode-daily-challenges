func arithmeticTriplets(nums []int, diff int) int {
    presence := make(map[int]bool)
    for _, n := range nums {
        presence[n] = true
    }

    result := 0
    diff2 := 2 * diff
    for _, num := range nums {
        if presence[num + diff] && presence[num + diff2] {
            result += 1
        }
    }
    return result
}