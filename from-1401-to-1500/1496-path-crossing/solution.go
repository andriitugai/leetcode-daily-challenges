type point struct {
    x int
    y int
}

func (p *point) move(dir rune) point {
    if dir == 'N' {
        return point{p.x, p.y + 1}
    } else if dir == 'W' {
        return point{p.x - 1, p.y}
    } else if dir == 'S' {
        return point{p.x, p.y - 1}
    } else if dir == 'E' {
        return point{p.x + 1, p.y}
    }
    return *p
}

func isPathCrossing(path string) bool {
    visited := make(map[point]bool)
    
    curr := point{x: 0, y: 0}
    visited[curr] = true

    for _, dir := range path {
        curr = curr.move(dir)
        if visited[curr] {
            return true
        }
        visited[curr] = true
    }
    return false
}