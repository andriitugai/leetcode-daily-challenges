func countGoodRectangles(rectangles [][]int) int {
    maxSize := 0
    maxCount := 0
    for _, r := range rectangles {
        size := r[0]
        if r[1] < size {
            size = r[1]
        }
        if size > maxSize {
            maxSize = size
            maxCount = 1
        } else if size == maxSize {
            maxCount += 1
        }
    }
    return maxCount
}