func numberOfPoints(nums [][]int) int {
    carSpace := make(map[int]bool)
    for _, car := range nums {
        for i := car[0]; i <= car[1]; i ++ {
            carSpace[i] = true
        }
    }
    return len(carSpace)
}