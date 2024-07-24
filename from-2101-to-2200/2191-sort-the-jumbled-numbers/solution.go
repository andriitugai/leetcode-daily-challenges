type mVal struct {
    index int
    mapped int
}

func sortJumbled(mapping []int, nums []int) []int {
    n := len(nums)
    mappedVals := make(map[int]mVal)
    for i := 0; i < n; i ++ {
        m := 0
        for _, c := range strconv.Itoa(nums[i]) {
            m = m * 10 + mapping[c - '0']
        }
        mappedVals[nums[i]] = mVal{mapped: m, index: i}
    }
    sort.Slice(nums, func(i, j int) bool {
        if mappedVals[nums[i]].mapped == mappedVals[nums[j]].mapped {
            return mappedVals[nums[i]].index < mappedVals[nums[j]].index
        }
        return mappedVals[nums[i]].mapped < mappedVals[nums[j]].mapped
    })

    return nums
}