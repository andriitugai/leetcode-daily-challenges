unc largestInteger(num int) int {
    digits := []int{}
    odd, even := []int{}, []int{}
    for num > 0 {
        d := num % 10
        digits = append(digits, d)
        if d % 2 == 0 {
            even = append(even, d)
        } else {
            odd = append(odd, d)
        }
        num /= 10
    }
    sort.Slice(odd, func(i, j int) bool {
        return odd[i] > odd[j]
    })
    sort.Slice(even, func(i, j int) bool {
        return even[i] > even[j]
    })

    result := 0
    po, pe := 0, 0
    var d int
    for i := len(digits) - 1; i >= 0; i-- {
        if digits[i] % 2 == 0 {
            d = even[pe]
            pe += 1
        } else {
            d = odd[po]
            po += 1
        }
        result = result * 10 + d
    }
    return result
}