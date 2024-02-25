func canTraverseAllPairs(nums []int) bool {
    n := len(nums)
    if n == 1 {
        return true
    }
    minVal, maxVal := nums[0], nums[0]
    for _, num := range nums {
        if num < minVal {
            minVal = num
        } else if num > maxVal {
            maxVal = num
        }
    }
    if minVal == 1 {
        return false
    }
    if minVal == maxVal {
        return true
    }

    v := make([][]int, 1 + maxVal)
    for i := 2; i <= maxVal; i++ {
        if v[i] == nil {
            for j := i; j <= maxVal; j += i {
                v[j] = append(v[j], i)
            }
        }
    }

    idx := make([]int, 1 + maxVal)
    for i := range idx {
        idx[i] = -1
    }
    // To solve the problem let's use union find
    items := make([]int, n)
    ranks := make([]int, n)
    for i := range items {
        items[i] = i
        ranks[i] = 1
    }
    // and keep tracking of the number of connected components
    cnt := n
    // Complexity: time: O(n*log(n)), space: O(n)
    for i, num := range nums {
        for _, p := range v[num] {
            if idx[p] == -1 {
                idx[p] = i
            } else if union(items, ranks, idx[p], i) {
                cnt -= 1
            }
        }
    }
    return cnt == 1
}

func sievePrimes(n int) []int {
    notPrimes := make([]int, n + 1)
    notPrimes[0], notPrimes[1] = 1, 1 

    for i := 0; i * i <= n; i ++ {
        if notPrimes[i] == 0 {
            for j := i * i; j <= n; j += i {
                notPrimes[j] = 1
            }
        }
    }
    primes := []int{}
    for i := 2; i <= n; i ++ {
        if notPrimes[i] == 0 {
            primes = append(primes, i)
        }
    }
    return primes
}

func find(items []int, i int) int {
    c, p := i, items[i]
    if c != p {
        p = find(items, p)
        items[i] = p
    }
    return p
}

func union(items, ranks []int, i, j int) bool {
    u := find(items, i)
    v := find(items, j)
    if u == v {
        return false
    }
    q := ranks[u] - ranks[v]
    if q > 0 {
        items[v] = u
    } else if q < 0 {
        items[u] = v
    } else {
        items[u] = v
        ranks[v] += 1
    }
    return true
}