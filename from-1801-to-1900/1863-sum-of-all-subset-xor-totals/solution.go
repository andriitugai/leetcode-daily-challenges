func subsetXORSum(nums []int) int {
    bits := 0
    for _, num := range nums {
        bits |= num
    }
    return bits * (1 << (len(nums) - 1))
}