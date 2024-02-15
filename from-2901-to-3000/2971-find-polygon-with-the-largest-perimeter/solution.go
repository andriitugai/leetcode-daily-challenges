func largestPerimeter(nums []int) int64 {
    var summ int64 = 0
    for _, num := range nums {
        summ += int64(num)
    }
    sort.Ints(nums)
    n := len(nums)
    for n >= 3 && int64(nums[n - 1]) >= summ - int64(nums[n - 1]){
        summ -= int64(nums[n - 1])
        n -= 1
    }
    if n < 3 {
        return int64(-1)
    }
    return summ
}