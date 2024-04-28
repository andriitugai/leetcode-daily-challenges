func addedInteger(nums1 []int, nums2 []int) int {
    mx1, mx2 := -1, -1
    for i := 0; i < len(nums1); i ++ {
        if nums1[i] > mx1 {
            mx1 = nums1[i]
        }
        if nums2[i] > mx2 {
            mx2 = nums2[i]
        }
    }
    return mx2 - mx1
}