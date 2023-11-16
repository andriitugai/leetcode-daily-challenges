unc findDifferentBinaryString(nums []string) string {
    n := len(nums[0])
    present := make(map[int64]bool)

    for _, bin := range nums {
        num, _ := strconv.ParseInt(bin, 2, 64)
        present[num] = true
    }

    var num int64 = 0
    for num <= 1 << n {
        if !present[num] {
            ans := strconv.FormatInt(num, 2)
            return strings.Repeat("0", n - len(ans)) + ans
        }
        num += 1
    }

    return "XXX"
}