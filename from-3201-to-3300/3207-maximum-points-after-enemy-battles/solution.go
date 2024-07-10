func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
    sort.Ints(enemyEnergies)
    var result int64 = 0
    left, right := 0, len(enemyEnergies) - 1
    for left <= right {
        if currentEnergy >= enemyEnergies[left] {
            result += int64(currentEnergy) / int64(enemyEnergies[left])
            currentEnergy = currentEnergy % enemyEnergies[left]
        } else {
            currentEnergy += enemyEnergies[right]
            right -= 1
        }
        if result == 0 {
            break
        }
    }
    return result
}