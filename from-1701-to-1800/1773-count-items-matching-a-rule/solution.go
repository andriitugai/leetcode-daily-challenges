func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    valIdx := 0
    if ruleKey == "color" {
        valIdx = 1
    } else if ruleKey == "name" {
        valIdx = 2
    }
    result := 0
    for _, item := range items {
        if ruleValue == item[valIdx] {
            result += 1
        }
    }
    return result
}