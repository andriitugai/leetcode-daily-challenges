func uniqueOccurrences(arr []int) bool {
    cnt := map[int]int{}
    for _, num := range arr {
        cnt[num] += 1
    }
    seen := make(map[int]bool)
    for _, occu := range cnt {
        if seen[occu] {
            return false
        }
        seen[occu] = true
    }
    return true
}