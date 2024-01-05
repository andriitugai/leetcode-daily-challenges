func moveZeroes(nums []int)  {
    ps, pf := 0, 0
    for ps < len(nums) && nums[ps] != 0 {
        ps += 1
    }
    pf = ps + 1
    for pf < len(nums) {
        if nums[pf] != 0 {
            nums[ps] = nums[pf]
            ps += 1
        }
        pf += 1
    }
    for ps < len(nums) {
        nums[ps] = 0
        ps += 1
    }
}
