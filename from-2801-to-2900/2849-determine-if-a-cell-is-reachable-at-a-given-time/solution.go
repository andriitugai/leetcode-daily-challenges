func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
    if sx == fx && sy == fy {
        return t != 1
    }
    dx := sx - fx
    if dx < 0 {
        dx = -dx
    }
    dy := sy - fy
    if dy < 0 {
        dy = -dy
    }

    return dx <= t && dy <= t
}