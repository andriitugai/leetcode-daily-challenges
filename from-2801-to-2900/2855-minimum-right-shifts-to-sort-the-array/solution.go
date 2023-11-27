func minimumRightShifts(nums []int) int {
    n := len(nums)
    srt := make([]int, n)
    copy(srt, nums)
    sort.Ints(srt)

    idx := 0
    for srt[0] != nums[idx] {
        idx += 1
        if idx == n {
            idx = 0
        }
    }

     for i := 1; i < n; i ++ {
         if srt[i] != nums[(idx + i) % n] {
             return -1
         }
     }
     if idx == 0 {
         return 0
     }
     return n - idx
}