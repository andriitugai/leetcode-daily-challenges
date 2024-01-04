func minOperations(nums []int) int {
    counts := make(map[int]int)
    for _, num := range nums {
        counts[num] += 1
    }

    result := 0
    for _, cnt := range counts {
        if cnt == 1 {
            return -1
        }

        for cnt >= 5 {
            cnt -= 3
            result += 1
        }
        if cnt == 2 || cnt == 3{
            result += 1
        } else {
            result += 2
        }
    }
    return result
}