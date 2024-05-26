func numberOfPairs(nums1 []int, nums2 []int, k int) int {
    result := 0
    for i := 0; i < len(nums1); i ++ {
        for j := 0; j < len(nums2); j ++ {
            if nums1[i] % (nums2[j] * k) == 0 {
                result += 1
            }
        }
    }
    return result
}