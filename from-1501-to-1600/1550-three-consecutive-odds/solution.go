func threeConsecutiveOdds(arr []int) bool {
    checkSum := 0
    for _, num := range arr {
        if num % 2 == 0 {
            checkSum = 0
        } else {
            checkSum += 1
            if checkSum == 3 {
                return true
            }
        }
    }
    return false
}