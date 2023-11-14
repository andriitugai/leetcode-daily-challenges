func pivotInteger(n int) int {
    x2 := (n * n + n) / 2
    sqRoot := int(math.Sqrt(float64(x2)))
    if sqRoot * sqRoot == x2 {
        return sqRoot
    }
    return -1
}