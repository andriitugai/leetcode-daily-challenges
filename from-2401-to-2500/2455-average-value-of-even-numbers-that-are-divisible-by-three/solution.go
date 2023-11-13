func averageValue(nums []int) int {
    total := 0
    n := 0

    for _, num := range nums {
        if num % 2 == 0 && num % 3 == 0 {
            total += num
            n += 1
        }
    }
    if n == 0 {
        return 0
    }
    return total / n
}