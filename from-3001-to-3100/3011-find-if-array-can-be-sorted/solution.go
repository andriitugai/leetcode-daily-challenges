unc canSortArray(nums []int) bool {
    sort.SliceStable(nums, func(i, j int) bool {
		return bits.OnesCount(uint(nums[i])) == bits.OnesCount(uint(nums[j])) && nums[i] < nums[j]
	})
	return slices.IsSorted(nums)
}