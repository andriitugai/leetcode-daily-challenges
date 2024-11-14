import (
    "math"
)

func minimizedMaximum(n int, quantities []int) int {
    low, high := 1, 0
    for _, quantity := range quantities {
        if quantity > high {
            high = quantity
        }
    }
    result := 0

    for low <= high {
        mid := low + (high-low)/2
        storesNeeded := 0
        for _, quantity := range quantities {
            storesNeeded += int(math.Ceil(float64(quantity) / float64(mid)))
        }
        
        if storesNeeded <= n {
            result = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    
    return result
}