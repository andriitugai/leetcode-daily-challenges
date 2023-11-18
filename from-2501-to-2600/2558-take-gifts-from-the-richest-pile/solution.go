func pickGifts(gifts []int, k int) int64 {
    var result int64 = 0
    for i := 0; i < len(gifts); i ++ {
        result += int64(gifts[i])
    }
    
    for t := 0; t < k; t ++ {
        maxIdx := -1
        maxGft := -1
        for i := 0; i < len(gifts); i ++ {
            if gifts[i] > maxGft {
                maxGft = gifts[i]
                maxIdx = i
            }
        }
        sqrt := int(math.Floor(math.Sqrt(float64(gifts[maxIdx]))))
        result -= int64(gifts[maxIdx] - sqrt)
        gifts[maxIdx] = sqrt
    }
    return result
}