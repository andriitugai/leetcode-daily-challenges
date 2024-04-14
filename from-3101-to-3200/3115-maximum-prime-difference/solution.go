func maximumPrimeDifference(nums []int) int {
    primes := make([]bool, 101)
    for i := 2; i < 101; i ++ {
        primes[i] = true
    }
    for p := 2; p * p <= 100; p ++ {
        if primes[p] {
            for i := p * p; i <= 100; i += p {
                primes[i] = false
            }
        }
    }

    leftPrimeIdx, rightPrimeIdx := -1, -1
    for i := 0; i < len(nums); i ++ {
        if primes[nums[i]] {
            leftPrimeIdx = i
            break
        }
    }
    for i := len(nums) - 1; i >= 0; i -- {
        if primes[nums[i]] {
            rightPrimeIdx = i
            break
        }
    }
    return rightPrimeIdx - leftPrimeIdx
}