func longestAlternatingSubarray(nums []int, threshold int) int {
    l, r := 0, 0
    maxLen := 0

    for r <= len(nums) - 1 {
        if nums[l] % 2 != 0 {
            l += 1
            r += 1
            continue
        }
        if nums[r] <= threshold {
           
           ln := (r - l) + 1     
            if ln > maxLen {
                maxLen = ln
            }   
        } else {
            r += 1
            l = r 
            continue
        }
        if (r < len(nums) - 1) && (nums[r] % 2 != nums[r + 1] % 2) {
            r += 1
        } else {
            r += 1
            l = r 
        }
    }
    return maxLen
}