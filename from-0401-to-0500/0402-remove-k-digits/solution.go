func removeKdigits(num string, k int) string {
    digits := strings.Split(num, "")
    stack := []string{}
    for i := 0; i < len(digits); i ++ {
        for k > 0 && len(stack) > 0 && stack[len(stack) - 1] > digits[i] {
            stack = stack[:len(stack) - 1]
            k -= 1
        }
        stack = append(stack, digits[i])
    }
    
    stack = stack[:len(stack) - k]

    lz := 0
    for lz < len(stack) && stack[lz] == "0" {
        lz += 1
    }
    stack = stack[lz:]
    if len(stack) > 0 {
        return strings.Join(stack, "")
    }
    return "0"
}