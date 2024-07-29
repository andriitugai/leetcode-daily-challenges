func nonSpecialCount(l int, r int) int {
    limit := int(math.Sqrt(float64(r)))
    sieve := make([]bool, limit + 1)
    sieve[0], sieve[1] = true, true

    for i := 2; i * i <= limit; i ++ {
        if !sieve[i] {
            for j := i * i; j <= limit; j += i {
                sieve[j] = true
            }
        }
    }

    spec := 0
    for i := 2; i <= limit; i ++ {
        if !sieve[i] && i * i <= r && i * i >= l {
            spec += 1
        }
    }

    total := r - l + 1
    return total - spec
}