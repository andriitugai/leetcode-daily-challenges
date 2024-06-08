func checkSubarraySum(nums []int, k int) bool {
    mods := make(map[int]int)
    mods[0] = 0
    s := 0
    for i, num := range nums {
        s += num
        mod := s % k
        if iPrev, ok := mods[mod]; !ok {
            mods[mod] = i + 1
        } else if iPrev < i {
            return true
        }
    }
    return false
}