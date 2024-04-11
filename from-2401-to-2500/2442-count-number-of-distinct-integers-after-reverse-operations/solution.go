func countDistinctIntegers(nums []int) int {
    reverse := func(num int) int {
        result := 0
        for num > 0 {
            result = result * 10 + num % 10
            num /= 10
        }
        return result
    }
    seen := make(map[int]bool)
    for _, num := range nums {
        seen[num] = true
        seen[reverse(num)] = true
    }
    return len(seen)
}