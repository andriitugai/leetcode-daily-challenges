func missingInteger(nums []int) int {
    count := nums[0]
    elems := make(map[int]bool)
    for i := 0; i < len(nums); i ++ {
        elems[nums[i]] = true
    }
    for i := 1; i < len(nums); i ++ {
        if nums[i] == nums[i - 1] + 1 {
            count += nums[i]
        } else {
            break
        }
    }
    for true {
        if !elems[count] {
            break
        } else {
            count += 1
        }
    }
    return count
}