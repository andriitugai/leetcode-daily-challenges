func captureForts(forts []int) int {
    max := 0
    empty, own := -1, -1
    for i := range forts {
        switch forts[i] {
            case -1:
                if own >= 0 && i - own - 1 > max {
                    max = i - own - 1
                }
                empty, own = i, -1
            case 1:
                if empty >= 0 && i - empty - 1 > max {
                    max = i - empty - 1
                }
                empty, own = -1, i
        }
    }
    return max
}