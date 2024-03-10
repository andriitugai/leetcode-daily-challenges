func intersection(nums1 []int, nums2 []int) []int {
    imap1 := make(map[int]bool)
    for _, num := range nums1{
        imap1[num] = true
    }
    imap2 := make(map[int]bool)
    for _, num := range nums2{
        imap2[num] = true
    }
    result := []int{}
    for num, _ := range imap1 {
        if imap2[num] {
            result = append(result, num)
        }
    }
    return result
}