func findErrorNums(nums []int) []int {
    cnt := make(map[int]int)
    var twice, missing int
    for _, num := range nums {
        cnt[num] += 1
        if cnt[num] == 2 {
            twice = num
        }
    }
    for num := 1; num <= len(nums); num ++ {
        if cnt[num] == 0 {
            missing = num
            break
        }
    }
    return []int{twice, missing}