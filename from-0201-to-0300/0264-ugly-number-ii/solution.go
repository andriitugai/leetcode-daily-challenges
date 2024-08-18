func nthUglyNumber(n int) int {
    uglies := []int{1}
    u2, u3, u5 := 0, 0, 0
    curr := 1

    for curr < n {
        newU := min3(uglies[u2] * 2, uglies[u3] * 3, uglies[u5] * 5)
        if newU == uglies[u2] * 2 {
            u2 += 1
        }
        if newU == uglies[u3] * 3 {
            u3 += 1
        }
        if newU == uglies[u5] * 5 {
            u5 += 1
        }
        uglies = append(uglies, newU)
        curr += 1
    }

    return uglies[n - 1]
}

func min3(a, b, c int) int {
    m3 := a 
    if b < m3 {
        m3 = b
    }
    if c < m3 {
        m3 = c
    }
    return m3
}