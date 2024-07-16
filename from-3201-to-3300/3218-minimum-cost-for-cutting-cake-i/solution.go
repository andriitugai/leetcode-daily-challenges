func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
    sort.Ints(horizontalCut)
    sort.Ints(verticalCut)
    sumH, sumV := 0, 0
    result := 0
    for _, cut := range horizontalCut {
        sumH += cut
    }
    for _, cut := range verticalCut {
        sumV += cut
    }

    row := len(horizontalCut) - 1
    col := len(verticalCut) - 1
    for row >= 0 && col >= 0 {
        if horizontalCut[row] > verticalCut[col] {
            result += horizontalCut[row] + sumV
            sumH -= horizontalCut[row]
            row -= 1
        } else {
            result += verticalCut[col] + sumH
            sumV -= verticalCut[col]
            col -= 1
        }
    }
    return result + sumH + sumV
}