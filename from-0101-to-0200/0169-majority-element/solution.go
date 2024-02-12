func majorityElement(nums []int) int {
    n := len(nums)
    counts := make(map[int]int)
    for _, num := range nums {
        counts[num] += 1
    }

    for num, cnt := range counts {
        if cnt > n / 2 {
            return num
        }
    }
    return -1
}