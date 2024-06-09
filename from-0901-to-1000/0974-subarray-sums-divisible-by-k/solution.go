func subarraysDivByK(nums []int, k int) int {
    mods := make(map[int]int)
    mods[0] = 1
    rem := 0
    result := 0
    for _, num := range nums {
        rem = (((rem + num) % k) + k) % k
        result += mods[rem]
        mods[rem] += 1
    }
    return result
}