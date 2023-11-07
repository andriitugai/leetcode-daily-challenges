func minMaxGame(nums []int) int {
    game := func() []int{
        newNums := make([]int, len(nums) / 2)
        for i := 0; i < len(nums) / 2; i ++ {
            newNums[i] = nums[2 * i]
            if i % 2 == 0 {
                if newNums[i] > nums[2 * i + 1] {
                    newNums[i] = nums[2 * i + 1]
                }
            } else {
                if newNums[i] < nums[2 * i + 1] {
                    newNums[i] = nums[2 * i + 1]
                }
            }
        }
        return newNums
    }
    for len(nums) > 1 {
        nums = game()
    }
    return nums[0]
}