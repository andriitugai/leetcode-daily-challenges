func evalRPN(tokens []string) int {
    stack := []int{}
    var op1, op2 int

    for _, token := range tokens {
        if token == "+" || token == "-" || token == "*" || token == "/" {
            op1 = stack[len(stack) - 1]
            op2 = stack[len(stack) - 2]
            stack = stack[:len(stack) - 2]
            if token == "+" {
                op1 = op1 + op2
            } else if token == "-" {
                op1 = op2 - op1
            } else if token == "*" {
                op1 = op1 * op2
            } else if token == "/" {
                op1 = op2 / op1
            } 
        } else {
            op1, _ = strconv.Atoi(token)
        }
        stack = append(stack, op1)
    }
    return stack[0]
}