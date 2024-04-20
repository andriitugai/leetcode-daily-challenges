func miceAndCheese(reward1 []int, reward2 []int, k int) int {
    n := len(reward1)
    diff := make([]int, n)
    total := 0
    for i := 0; i < n; i ++ {
        diff[i] = reward1[i] - reward2[i]
        total += reward2[i]
    }
    sort.Slice(diff, func(i, j int)bool {
        return diff[i] > diff[j]
    })

    for i := 0; i < k; i ++ {
        total += diff[i]
    }
    return total
}