func findMaxK(nums []int) int {
    hash := make(map[int]bool)
    ans := 0
    for _, num := range nums {
        if hash[-num] && abs(num) > ans {
            ans = abs(num)
        }
        hash[num] = true
    }
    if ans == 0 {
        return -1
    }
    return ans
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}