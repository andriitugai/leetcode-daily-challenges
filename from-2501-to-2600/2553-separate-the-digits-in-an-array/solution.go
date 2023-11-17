func separateDigits(nums []int) []int {
    result := []int{}
    for _, num := range nums {
        stack := []int{}
        for num > 0 {
            stack = append(stack, num % 10)
            num /= 10
        }
        for len(stack) > 0 {
            result = append(result, stack[len(stack)-1])
            stack = stack[:len(stack)-1]
        }
    }
    return result
}