func minPatches(nums []int, n int) int {
    curr := 1 //able to form numbers 1 to curr - 1, checking if curr is also possible
    added := 0
    i := 0

    for curr <= n {
        if (i < len(nums) && nums[i] <= curr) {
            curr += nums[i] //if able to form 1 to curr-1, and nums[i] <= curr, then can make 1 to (curr-1 + nums[i])
            i++
        } else { //if nums[i] > curr, then can't make curr, so have to add curr, then can make 1 to (curr - 1 + curr)
            added++
            curr += curr
        }
    }

    return added
}