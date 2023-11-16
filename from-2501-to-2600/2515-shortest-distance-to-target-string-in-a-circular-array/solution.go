func closetTarget(words []string, target string, startIndex int) int {
    n := len(words)
    idx := startIndex

    movesRight := 0
    for movesRight <= n {
        if words[idx] == target {
            break
        }
        idx += 1
        if idx == n {
            idx = 0
        } 
        movesRight += 1
    }
    if movesRight == n + 1 {
        return -1
    }

    idx = startIndex
    movesLeft := 0
    for movesLeft <= n {
        if words[idx] == target {
            break
        }
        idx -= 1
        if idx == -1 {
            idx = n - 1
        }
        movesLeft += 1
    }

    if movesLeft < movesRight {
        return movesLeft
    }
    return movesRight
}