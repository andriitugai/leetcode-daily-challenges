func numberOfBeams(bank []string) int {
    numLasersInRow := func(s string) int {
        n := 0
        for _, b := range s {
            if b == '1' {
                n += 1
            }
        }
        return n
    }
    prev := 0
    result := 0
    for r := len(bank) - 1; r >= 0; r -- {
        curr := numLasersInRow(bank[r])
        if curr > 0 {
            result += prev * curr
            prev = curr
        }
    }
    return result
}