func countSymmetricIntegers(low int, high int) int {
    result := 0
    for c := 11; c <= 99; c += 11 {
        if c >= low && c <= high {
            result += 1
        }
    }
    c := 1000
    for c <= high {
        if c >= low {
            if (c / 1000) + ((c /100) % 10) == c % 10 + (c % 100) / 10 {
                result += 1
            }
        }
        c += 1
    }
    return result
}