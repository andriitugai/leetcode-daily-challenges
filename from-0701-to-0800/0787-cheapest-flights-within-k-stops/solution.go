func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    prev := []int{}
    now := make([]int, n)
    
    for i := range now {
        now[i] = math.MaxInt32
    }
    
    now[src] = 0
    for k >= 0 {
        prev = now
        now = append([]int{}, prev...)
        
        for _, flight := range flights {
            now[flight[1]] = min(now[flight[1]], prev[flight[0]]+flight[2])
        }
        k -= 1
    }
    
    if now[dst] == math.MaxInt32 {
        return -1
    }
    
    return now[dst]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}