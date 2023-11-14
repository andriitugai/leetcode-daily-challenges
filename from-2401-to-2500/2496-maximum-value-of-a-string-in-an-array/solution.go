func maximumValue(strs []string) int {
    maxVal := -1
    for _, str := range strs {
        if val, err := strconv.Atoi(str); err == nil {
            if val > maxVal {
                maxVal = val
            }
        } else {
            if len(str) > maxVal {
                maxVal = len(str)
            }
        }
    }
    return maxVal
}