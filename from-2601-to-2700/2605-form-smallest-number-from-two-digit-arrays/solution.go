func minNumber(nums1 []int, nums2 []int) int {
    sort.Ints(nums1)
    sort.Ints(nums2)

    p1, p2 := 0, 0
    for p1 < len(nums1) && p2 < len(nums2) {
        if nums1[p1] == nums2[p2] {
            return nums1[p1]
        } else if nums1[p1] < nums2[p2] {
            p1 += 1
        } else {
            p2 += 1
        }
    }

    if nums1[0] < nums2[0] {
        return nums1[0] * 10 + nums2[0]
    }
    return nums2[0] * 10 + nums1[0]
}