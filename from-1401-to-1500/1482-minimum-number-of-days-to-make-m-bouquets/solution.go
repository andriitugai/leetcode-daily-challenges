func minDays(bloomDay []int, m int, k int) int {
    if len(bloomDay) < m * k {
        return -1
    }
    left, right := bloomDay[0], bloomDay[0]
    for i := 1; i < len(bloomDay); i ++ {
        if bloomDay[i] < left {
            left = bloomDay[i]
        } else if bloomDay[i] > right {
            right = bloomDay[i]
        }
    }

    isPossible := func(day int) bool {
        bouquets := 0
        adj := 0
        for _, bDay := range bloomDay {
            if bDay > day {
                adj = 0
            } else {
                adj += 1
                if adj == k {
                    bouquets += 1
                    if bouquets == m {
                        return true
                    }
                    adj = 0
                }
            }
        }
        return false
    }

    for left < right {
        mid := left + (right - left) / 2
        if isPossible(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}