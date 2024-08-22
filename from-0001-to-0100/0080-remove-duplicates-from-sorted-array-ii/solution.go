func removeDuplicates(nums []int) int {
    nDupl := 1
    prev := nums[0]
    tail := 1
    for i := 1; i < len(nums); i ++ {
        if nums[i] == prev {
            nDupl += 1
            if nDupl <= 2 {
                nums[tail] = nums[i]
                tail += 1
            }
        } else {
            nDupl = 1
            nums[tail] = nums[i]
            tail += 1
        }
        prev = nums[i]
    }
    return tail
}