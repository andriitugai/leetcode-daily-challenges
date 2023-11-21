func distinctDifferenceArray(nums []int) []int {
    result := make([]int, len(nums))
    mPre, mSuf := make(map[int]int), make(map[int]int)
    for _, num := range nums {
        mSuf[num] += 1
    }
    for i := 0; i < len(nums); i ++ {
        mPre[nums[i]] += 1
        mSuf[nums[i]] -= 1
        if mSuf[nums[i]] == 0 {
            delete(mSuf, nums[i])
        }
        result[i] = len(mPre) - len(mSuf)
    }
    return result
}