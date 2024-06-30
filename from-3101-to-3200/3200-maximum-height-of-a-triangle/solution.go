func maxHeightOfTriangle(red int, blue int) int {
    redFirst := helper(red, blue)
    blueFirst := helper(blue, red)
    if redFirst > blueFirst {
        return redFirst
    }
    return blueFirst
}

func helper(red int, blue int) int {
    curHeight := 0
    rowToComplete := 1
    for true {
        if rowToComplete % 2 == 1 {
            if red < rowToComplete {
                break
            }
            red -= rowToComplete
        } else {
            if blue < rowToComplete {
                break
            }
            blue -= rowToComplete
        }
        curHeight += 1
        rowToComplete += 1
    }
    return curHeight
}