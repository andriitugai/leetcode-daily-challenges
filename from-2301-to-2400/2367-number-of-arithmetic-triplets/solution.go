func arithmeticTriplets(nums []int, diff int) int {
    presence := make(map[int]bool)
    result := 0
    for _, num := range nums {
        presence[num] = true
        if num >= diff * 2 && presence[num - diff] && presence[num - diff * 2] {
            result += 1
        }
    }
    return result
}