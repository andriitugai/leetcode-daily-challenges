type point struct {
    dist int
    tag byte
}

func maxPointsInsideSquare(points [][]int, s string) int {
    n := len(points)
    pts := make([]point, n)
    for i := 0; i < n; i ++ {
        pts[i] = point{dist: maxAbs(points[i][0], points[i][1]), tag: s[i],}
    }

    sort.Slice(pts, func(i, j int) bool {
        return pts[i].dist < pts[j].dist
    })
    
    result := 0
    idx := 0
    tags := make(map[byte]bool)

    for idx < n {
        curDist := pts[idx].dist
        for i := idx; i < n && pts[i].dist == curDist; i ++ {
            if tags[pts[i].tag] {
                return result
            }
            tags[pts[i].tag] = true
            idx += 1
        }
        result = idx
    }
    return result
}

func maxAbs(a, b int) int {
    if a < 0 {
        a = -a
    }
    if b < 0 {
        b = -b
    }
    if a > b {
        return a
    }
    return b
}