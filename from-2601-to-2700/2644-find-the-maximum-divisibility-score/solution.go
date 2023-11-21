func maxDivScore(nums []int, divisors []int) int {
    scores := make(map[int]int)
    for _, div := range divisors {
        score := 0
        for _, num := range nums {
            if num % div == 0 {
                score += 1
            }
        }
        scores[div] = score
    }
    maxScore := -1
    maxScoreDiv := -1
    for _, div := range divisors {
        if scores[div] > maxScore {
            maxScore = scores[div]
            maxScoreDiv = div
        } else if scores[div] == maxScore && div < maxScoreDiv {
            maxScoreDiv = div
        }
    }
    return maxScoreDiv
}