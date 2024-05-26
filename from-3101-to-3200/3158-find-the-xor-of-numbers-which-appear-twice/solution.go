func duplicateNumbersXOR(nums []int) int {
    result := 0
    counter := make(map[int]int)
    for _, num := range nums {
        counter[num] += 1
        if counter[num] == 2 {
            result ^= num
        }
    }
    return result
}