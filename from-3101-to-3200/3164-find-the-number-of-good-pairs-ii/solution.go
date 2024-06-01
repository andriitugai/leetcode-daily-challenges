func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
    divCount := make(map[int]int)
    for _, num := range nums2 {
        divCount[num * k] += 1
    }

    var result int64
    for _, num := range nums1 {
        for j := 1; j * j <= num; j ++ {
            if num % j == 0 {
                if dCount, ok := divCount[j]; ok {
                    result += int64(dCount)
                }
                ja := num / j
                if j != ja {
                    if daCount, ok := divCount[ja]; ok {
                        result += int64(daCount)
                    }
                }
            }
        }
    }
    return result
}