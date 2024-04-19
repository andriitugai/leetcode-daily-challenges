func wiggleSort(nums []int)  {
    sort.Ints(nums)
    cp := make([]int, len(nums))
    copy(cp, nums)
    right := len(nums) - 1
    mid := right / 2
    for i := 0; i < len(nums); i ++ {
        if i % 2 == 0 {
            nums[i] = cp[mid]
            mid --
        } else {
            nums[i] = cp[right]
            right --
        }
    }
}