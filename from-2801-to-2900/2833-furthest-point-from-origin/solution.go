func furthestDistanceFromOrigin(moves string) int {
    ml, mr, m_ := 0, 0, 0
    for i := 0; i < len(moves); i ++ {
        if moves[i] == 'L' {
            ml += 1
        } else if moves[i] == 'R' {
            mr += 1
        } else {
            m_ += 1
        }
    }
    if ml > mr {
        return ml + m_ - mr
    }
    return mr + m_ - ml
}