func maximumSetSize(nums1 []int, nums2 []int) int {
    n := len(nums1)
    set1 := make(map[int]bool)
    set2 := make(map[int]bool)

    for i := 0; i < n; i ++ {
        set1[nums1[i]] = true
        set2[nums2[i]] = true
    }

    half := n / 2
    for num := range set1 {
        if len(set1) <= half {
            break
        }
        if set2[num] {
            delete(set1, num)
        }
    }
    for num := range set1 {
        if len(set1) > half {
            delete(set1, num)
        } else {
            break
        }
    }
    result := len(set1)
    count := 0

    for num := range set2 {
        if !set1[num] && count < half {
            result += 1
            count += 1
        }
    }
    return result
}