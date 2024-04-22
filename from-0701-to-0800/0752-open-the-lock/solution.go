func openLock(deadends []string, target string) int {
    seen := make(map[string]bool)
    for _, de := range deadends {
        seen[de] = true
    }
    if seen["0000"] {
        return -1
    }

    neighbors := func(curr string) []string {
        result := []string{}
        nums, nei := make([]int, 4), make([]int, 4)
        var neiNum int
        var neiStr string
        for i := 0; i < 4; i ++ {
            nums[i], _ = strconv.Atoi(string(curr[i]))
        }
        for i := 0; i < 4; i ++ {
            copy(nei, nums)
            nei[i] = (nums[i] + 1) % 10
            neiNum = 0
            for j := 0; j < 4; j ++ {
                neiNum = neiNum * 10 + nei[j]
            }
            neiStr = fmt.Sprintf("%04d", neiNum)
            if !seen[neiStr] {
                result = append(result, neiStr)
            }

            nei[i] = (nums[i] + 9) % 10
            neiNum = 0
            for j := 0; j < 4; j ++ {
                neiNum = neiNum * 10 + nei[j]
            }
            neiStr = fmt.Sprintf("%04d", neiNum)
            if !seen[neiStr] {
                result = append(result, neiStr)
            }
        }
        return result
    }
    q := []string{"0000"}
    level := 0
    var size int
    var lock string
    for len(q) > 0 {
        size = len(q)
        for i := 0; i < size; i ++ {
            lock = q[0]
            if lock == target {
                return level
            }
            q = q[1:]
            seen[lock] = true
            for _, nei := range neighbors(lock) {
                seen[nei] = true
                q = append(q, nei)
            }
        }
        level += 1
    }
    return -1
}