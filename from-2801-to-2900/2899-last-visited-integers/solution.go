func lastVisitedIntegers(words []string) []int {
    result := []int{}
    stack := []int{}
    cnt := 0
    for _,item := range(words) {
        if item == "prev" {
            cnt += 1
            if cnt > len(stack){
                result = append(result, -1)
            } else {
                result = append(result, stack[len(stack) - cnt])
            }
        } else {
            val, _ := strconv.Atoi(item)
            stack = append(stack, val)
            cnt = 0
        }
    }
    return result
}