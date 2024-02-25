func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
    n := len(bottomLeft)
    maxSide := 0

    for i := 0; i < n; i ++ {
        for j := i + 1; j < n; j ++ {
            blx := max(bottomLeft[i][0], bottomLeft[j][0])
            bly := max(bottomLeft[i][1], bottomLeft[j][1])
            trx := min(topRight[i][0], topRight[j][0])
            try := min(topRight[i][1], topRight[j][1])

            if blx < trx && bly < try {
                side := min(trx - blx, try - bly)
                if side > maxSide {
                    maxSide = side
                }
            }
        }
    }
    return int64(maxSide) * int64(maxSide)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}