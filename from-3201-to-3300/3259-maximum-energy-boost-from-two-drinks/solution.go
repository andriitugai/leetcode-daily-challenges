func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
    n := len(energyDrinkA)
    var maxA, maxB int64 = 0, 0
    var newMaxA, newMaxB int64

    for i := 0; i < n; i ++ {
        newMaxA = max(maxA + int64(energyDrinkA[i]), maxB)
        newMaxB = max(maxB + int64(energyDrinkB[i]), maxA)
        maxA = newMaxA
        maxB = newMaxB
    }
    return max(maxA, maxB)
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}