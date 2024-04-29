func minimumAddedInteger(nums1 []int, nums2 []int) int {
    sort.Ints(nums1)
    sort.Ints(nums2)

    check := func(delta int) bool {
        missed := 0
        i2 := 0
        for i1 := 0; i1 < len(nums1) && i2 < len(nums2); i1 ++ {
            if nums1[i1] + delta != nums2[i2] {
                missed += 1
            } else {
                i2 += 1
            }
        }
        return missed <= 2
    }

    result := 2000
    for k := 0; k < 3; k ++ {
        diff := nums2[0] - nums1[k]
        if check(diff) && diff < result {
            result = diff
        }
    }
    return result
}