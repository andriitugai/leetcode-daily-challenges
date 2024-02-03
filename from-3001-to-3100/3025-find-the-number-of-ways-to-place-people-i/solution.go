func numberOfPairs(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        if points[i][1] == points[j][1] {
            return points[i][0] < points[j][0]
        }
        return points[i][1] > points[j][1]
    })
        
    cnt := 0
    
    for i := 0; i < len(points); i++ {
        prevDiff := math.MaxInt
        for j := i + 1; j < len(points); j++ {
            diff := points[j][0] - points[i][0]
            if diff >= 0 && diff < prevDiff {
                prevDiff = diff
                cnt++
            }
        }
    }

    return cnt
}