type point struct {
    x int
    y int
}

func robotSim(commands []int, obstacles [][]int) int {
    isObstacle := make(map[point]bool)
    for _, obs := range obstacles {
        isObstacle[point{obs[0], obs[1]}] = true
    }
    maxDist, curDist := 0, 0
    dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    dirIdx := 0
    x, y := 0, 0

    for _, c := range commands {
        if c == -1 {
            dirIdx = (dirIdx + 1) % 4
        } else if c == -2 {
            dirIdx = (dirIdx + 3) % 4
        } else {
            for i := 0; i < c; i ++ {
                newX := x + dirs[dirIdx][0]
                newY := y + dirs[dirIdx][1]
                if isObstacle[point{newX, newY}] {
                    break
                }
                x, y = newX, newY
                curDist = x * x + y * y
                if maxDist < curDist {
                    maxDist = curDist
                }
            }
        }
    }
    return maxDist
}