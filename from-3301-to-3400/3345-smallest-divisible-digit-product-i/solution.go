func smallestNumber(n int, t int) int {
    getProduct := func(n int) int {
        p := 1
        for n > 1 {
            p *= n % 10
            n /= 10
        }
        return p
    }
    for c := 0; c < t; c ++ {
        if getProduct(n + c) % t == 0 {
            return n + c
        }
    }
    return -1
}