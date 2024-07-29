func canAliceWin(nums []int) bool {
    s1, s2 := 0, 0
    for _, num := range nums {
        if num < 10 {
            s1 += num
        } else {
            s2 += num
        }
    }
    return s1 != s2
}