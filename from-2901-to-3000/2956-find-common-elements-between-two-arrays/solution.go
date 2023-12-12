func findIntersectionValues(nums1 []int, nums2 []int) []int {
    s1, s2 := make(map[int]bool), make(map[int]bool)
    for _, num := range nums1 {
        s1[num] = true
    }
    for _, num := range nums2 {
        s2[num] = true
    }
    r1 := 0
    for _, num := range nums1 {
        if s2[num] {
            r1 += 1
        }
    }
    r2 := 0
    for _, num := range nums2 {
        if s1[num] {
            r2 += 1
        }
    }

    return []int{r1, r2}
}