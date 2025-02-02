func maxDifference(s string) int {
    freq := make(map[rune]int)
    for _, c := range s {
        freq[c] += 1
    }

    maxOdd := 0
    minEven := len(s) + 1
    for _, f := range freq {
        if f % 2 == 1 && f > maxOdd {
            maxOdd = f
        } else if f % 2 == 0 && f < minEven {
            minEven = f
        }
    }
    return maxOdd - minEven
}