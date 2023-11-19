func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    p1, p2 := 0, 0
    result := [][]int{}
    for p1 < len(nums1) && p2 < len(nums2) {
        if nums1[p1][0] == nums2[p2][0] {
            result = append(result, []int{nums1[p1][0], nums1[p1][1] + nums2[p2][1]})
            p1 += 1
            p2 += 1
        } else if nums1[p1][0] < nums2[p2][0] {
            result = append(result, nums1[p1])
            p1 += 1
        } else {
            result = append(result, nums2[p2])
            p2 += 1
        }
    }
    for p1 < len(nums1) {
        result = append(result, nums1[p1])
        p1 += 1
    }
    for p2 < len(nums2) {
        result = append(result, nums2[p2])
        p2 += 1
    }
    return result
}