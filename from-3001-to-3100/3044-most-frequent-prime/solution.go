func mostFrequentPrime(mat [][]int) int {
    isPrime := func(num int) bool {
        if num < 2 {
            return false
        }
        if num == 2 {
            return true
        }
        for i := 2; i * i <= num; i ++ {
            if num % i == 0 {
                return false
            }
        }
        return true
    }
    dirs := [][]int{ 
        []int{0, 1}, []int{1, 1}, []int{1, 0}, []int{1, -1}, []int{0, -1}, []int{-1, -1}, []int{-1, 0}, []int{-1, 1},
    }
    countPrimes := make(map[int]int)
    m, n  := len(mat), len(mat[0])
    for r := 0; r < m; r ++ {
        for c := 0; c < n; c ++ {
            for _, dir := range dirs {
                num := 0
                dr, dc := dir[0], dir[1]
                cr, cc := r, c
                for cr >= 0 && cr < m && cc >= 0 && cc < n {
                    num = num * 10 + mat[cr][cc]
                    cr += dr
                    cc += dc
                    if num > 10 && isPrime(num) {
                        countPrimes[num] += 1
                    }
                }
            }
        }
    }
    maxFreqPrime := -1
    maxFreq := 0
    for prime, freq := range countPrimes {
        if freq > maxFreq {
            maxFreq = freq
            maxFreqPrime = prime
        } else if freq == maxFreq && prime > maxFreqPrime {
            maxFreqPrime = prime
        }
    }
    return maxFreqPrime
}