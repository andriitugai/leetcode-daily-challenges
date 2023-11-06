func digitSum(s string, k int) string {
    for len(s) > k {
        digits := strings.Split(s, "")
        sums := []int{}
        sum, cnt, i := 0, 0, 0
        for i < len(digits) {
            val, err := strconv.Atoi(digits[i])
            if err != nil { fmt.Println(err) }
            sum += val
            cnt++
            if cnt == k || i == len(digits) - 1 {
                sums = append(sums, sum)
                sum = 0
                cnt = 0
            }
            i++
        }
        s = ""
        for _, v := range sums { s += strconv.Itoa(v) }
    }
    return s
}