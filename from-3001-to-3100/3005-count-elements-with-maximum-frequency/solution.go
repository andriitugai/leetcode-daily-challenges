func maxFrequencyElements(nums []int) int {
    cnt := make(map[int]int)
    maxFreq := 0
    for _, num := range nums {
        cnt[num] += 1
        if cnt[num] > maxFreq {
            maxFreq = cnt[num]
        }
    }

    result := 0
    for _, freq := range cnt {
        if freq == maxFreq {
            result += maxFreq
        }
    }
    return result
}