func queryResults(limit int, queries [][]int) []int {
    colorMap := make(map[int]int)
    colorCount := make(map[int]int)
    result := make([]int, len(queries))
    numColors := 0

    for i, q := range queries {
        ball, newColor := q[0], q[1]
        if oldColor, ok := colorMap[ball]; ok {
            colorCount[oldColor] -= 1
            if colorCount[oldColor] == 0 {
                numColors -= 1
            }
        }

        colorMap[ball] = newColor
        colorCount[newColor] += 1
        if colorCount[newColor] == 1 {
            numColors += 1
        }

        result[i] = numColors
    }
    return result
}