func findClosest(x int, y int, z int) int {
    d1, d2 := z - x, z - y
    if d1 < 0 { d1 = -d1 }
    if d2 < 0 { d2 = -d2 }
    if d1 < d2 {
        return 1
    } else if d1 > d2 {
        return 2
    }
    return 0
}