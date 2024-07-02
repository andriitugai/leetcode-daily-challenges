func intersect(nums1 []int, nums2 []int) []int {
    c1, c2 := make(map[int]int), make(map[int]int)
    for _, num := range nums1 {
        c1[num] += 1
    }
    for _, num := range nums2 {
        c2[num] += 1
    }

    result := []int{}
    for k, v1 := range c1 {
        v2 := c2[k]
        for i := 0; i < min(v1, v2); i ++ {
            result = append(result, k)
        }
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}