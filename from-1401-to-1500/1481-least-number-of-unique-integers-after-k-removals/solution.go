func findLeastNumOfUniqueInts(arr []int, k int) int {
    counter := make(map[int]int)
    for _, num := range arr {
        counter[num] += 1
    }
    counts := []int{}
    for _, v := range counter {
        counts = append(counts, v)
    }
    sort.Ints(counts)

    cp := 0
    for cp < len(counts) && k >= counts[cp] {
        k -= counts[cp]
        cp += 1
    }
    return len(counts) - cp
}