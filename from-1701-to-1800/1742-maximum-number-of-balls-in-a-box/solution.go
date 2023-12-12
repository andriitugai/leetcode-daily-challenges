func countBalls(lowLimit int, highLimit int) int {
    cnt := make(map[int]int)
    getSumDigits := func(num int) int {
        result := 0
        for num > 0 {
            result += num % 10
            num /= 10
        }
        return result
    }
    for num := lowLimit; num <= highLimit; num ++ {
        cnt[getSumDigits(num)] += 1
    }
    maxBalls := -1
    for _, balls := range cnt {
        if balls > maxBalls {
            maxBalls = balls
        }
    }
    return maxBalls
}