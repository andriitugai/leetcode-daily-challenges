func nextPermutation(nums []int)  {
    i := len(nums) - 2
    for i > -1 {
        minToRight, swapIndex := nums[i+1], i+1
        for j := i+1; j <= len(nums) - 1; j++ { 
            if nums[j] < minToRight && nums[j] > nums[i] { 
                minToRight, swapIndex = nums[j], j
            }
        }
        if minToRight > nums[i] { 
            nums[swapIndex] = nums[i]; nums[i] = minToRight; break 
        }
        i--
  }
  sort.Ints(nums[i+1:]) 
}