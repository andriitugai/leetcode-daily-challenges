func destroyTargets(nums []int, space int) int {
    sort.Ints(nums)
    mods := make(map[int]int)
    maxModCount := 0
    var md int
    for _, num := range nums {
        md = num % space
        mods[md] += 1
        if mods[md] > maxModCount {
            maxModCount = mods[md]
        }
    }
    for i := 0; i < len(nums); i ++ {
        if mods[nums[i] % space] == maxModCount {
            return nums[i]
        }
    }
    return -1
}