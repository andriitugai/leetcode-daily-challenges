type info struct {
    val int
    freq int
}

func frequencySort(nums []int) []int {
    n := len(nums)
    cnt := make(map[int]int)
    for _, num := range nums {
        cnt[num] += 1
    }
    freqs := make([]info, n)
    for i := 0; i < n; i ++ {
        freqs[i] = info{val: nums[i], freq: cnt[nums[i]]}
    }
    sort.Slice(freqs, func(i, j int) bool {
        if freqs[i].freq == freqs[j].freq {
            return freqs[i].val > freqs[j].val
        }
        return freqs[i].freq < freqs[j].freq
    })
    result := make([]int, n)
    for i := 0; i < n; i ++ {
        result[i] = freqs[i].val
    }
    return result
}