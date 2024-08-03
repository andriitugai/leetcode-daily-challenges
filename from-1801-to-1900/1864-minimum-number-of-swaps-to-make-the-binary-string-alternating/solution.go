func minSwaps(s string) int {
    n := len(s)
    oddOnes, oddZeros, evenOnes, evenZeros := 0, 0, 0, 0
    for i := 0; i < n; i ++ {
        if i % 2 == 0 {
            if s[i] == '1' {
                evenOnes += 1
            } else {
                evenZeros += 1
            }
        } else {
            if s[i] == '1' {
                oddOnes += 1
            } else {
                oddZeros += 1
            }
        }
    }
    if abs(oddOnes + evenOnes - oddZeros - evenZeros) > 1 {
        return -1
    }
    if n % 2 == 0 {
        if oddOnes == evenZeros && (evenOnes == 0 || oddOnes == 0){
            return 0
        } else {
            return min(abs(n / 2 - oddOnes), abs(n / 2 - evenOnes))
        }
    } else {
        if oddOnes + evenOnes > oddZeros + evenZeros {
            return evenZeros
        } 
    }
    return evenOnes
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}