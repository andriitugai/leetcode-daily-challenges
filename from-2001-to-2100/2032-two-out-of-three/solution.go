func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    cnt := make(map[int]int)
    for _, num := range nums1 {
        cnt[num] |= 1
    }
    for _, num := range nums2 {
        cnt[num] |= 2
    }
    for _, num := range nums3 {
        cnt[num] |= 4
    }

    result := []int{}
    for k, v := range cnt {
        if v == 3 || v == 5 || v == 6 || v == 7{
            result = append(result, k)
        }
    }

    return result
}