func bagOfTokensScore(tokens []int, power int) int {
    sort.Ints(tokens)
    score, maxScore := 0, 0

    for len(tokens) > 0 {
        if power >= tokens[0] {
            power -= tokens[0]
            tokens = tokens[1:]
            score += 1
            if score > maxScore {
                maxScore = score
            }
        } else if score > 0 {
            power += tokens[len(tokens) - 1]
            tokens = tokens[:len(tokens) - 1]
            score -= 1
        } else {
            break
        }
    }
    return maxScore
}