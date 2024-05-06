func minAnagramLength(s string) int {
    totLen := len(s)
    cnt := make(map[byte]int)
    for i := 0; i < len(s); i ++ {
        cnt[s[i]] += 1
    }
    counts := []int{}
    for _, v := range cnt {
        counts = append(counts, v)
    }
    divisors := getAllCommonDivisors(counts)
    // divisors are sorted in descending order
    for _, div := range divisors {
        if totLen % div != 0 {
            continue
        }
        // check for anagrams
        isConcat := true
        for i := 0; i < div; i ++ {
            anaCnt := make(map[byte]int)
            anaLen := totLen / div
            for j := 0; j < anaLen; j ++ {
                anaCnt[s[i * anaLen + j]] += 1
            }
            isPart := true
            for k, v := range cnt {
                if anaCnt[k] * div != v {
                    isPart = false
                    break
                }
            }
            if !isPart {
                isConcat = false
                break
            }
        }
        if isConcat {
            return totLen / div
        }
    }
    return totLen
}

func gcd(a, b int) int {
    if a == 0 {
        return b
    }
    return gcd(b % a, a)
}

func getAllCommonDivisors(arr []int) []int {
    g := arr[0]
    for _, num := range(arr[1:]) {
        g = gcd(g, num)
    }

    divisors := []int{}
    for i := 1; i * i <= g; i ++ {
        if g % i == 0 {
            divisors = append(divisors, i)
            if i * i != g {
                divisors = append(divisors, g / i)
            }
        }
    }
    sort.Slice(divisors, func(i, j int) bool {
        return divisors[i] > divisors[j]
    })
    return divisors
}