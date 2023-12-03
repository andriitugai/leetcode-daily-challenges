func maximumUnits(boxTypes [][]int, truckSize int) int {
    sort.Slice(boxTypes, func(i, j int) bool {
        return boxTypes[i][1] > boxTypes[j][1]
    })
    result := 0
    for i := 0; i < len(boxTypes); i ++ {
        numBoxes := boxTypes[i][0]
        numUnitsPerBox := boxTypes[i][1]
        if numBoxes > truckSize {
            result += truckSize * numUnitsPerBox
            break
        }
        result += numBoxes * numUnitsPerBox
        truckSize -= numBoxes
    }
    return result
}