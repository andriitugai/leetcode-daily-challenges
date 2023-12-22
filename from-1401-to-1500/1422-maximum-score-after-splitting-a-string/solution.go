func maxScore(s string) int {
    nOnes := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '1' {
            nOnes += 1
        }
    }

    score, maxScore := nOnes, 0
    for i := 0; i < len(s)-1; i++ {
        if s[i] == '0' {
            score += 1
        } else {
            score -= 1
        }
        if score > maxScore {
            maxScore = score
        }
    }
    return maxScore
}