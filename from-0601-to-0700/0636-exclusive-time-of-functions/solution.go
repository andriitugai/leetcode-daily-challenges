func exclusiveTime(n int, logs []string) []int {
    result := make([]int, n)
    stack := [][]int{}

    for _, log := range logs {
        data := strings.Split(log, ":")
        id, _ := strconv.Atoi(data[0])
        event := data[1]
        time, _ := strconv.Atoi(data[2])
        if event == "start" {
            if len(stack) > 0 {
                result[stack[len(stack)-1][0]] += (time - stack[len(stack) - 1][1])
            }
            stack = append(stack, []int{id, time})
        } else {
            popped := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            result[popped[0]] += time - popped[1] + 1
            if len(stack) > 0 {
                stack[len(stack) - 1][1] = time + 1
            }
        }
    }
    return result
}