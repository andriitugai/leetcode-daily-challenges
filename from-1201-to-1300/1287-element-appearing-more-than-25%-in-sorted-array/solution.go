func findSpecialInteger(arr []int) int {
    goal := len(arr) / 4
    curr := -99
    cnt := 0

    for _, num := range arr {
        if num != curr {
            curr = num
            cnt = 0
        }
        cnt += 1
        if cnt > goal {
            return curr
        }
    }
    return -99
}
