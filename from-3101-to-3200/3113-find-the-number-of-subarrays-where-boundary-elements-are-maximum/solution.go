func numberOfSubarrays(nums []int) int64 {
    var result int64 = 0
    stack := [][]int{}
    for _, num := range nums {
        for len(stack) > 0 && stack[len(stack) - 1][0] < num {
            stack = stack[:len(stack) - 1]
        }
        if len(stack) == 0 || stack[len(stack) - 1][0] > num {
            stack = append(stack, []int{num, 0})
        }
        stack[len(stack) - 1][1] += 1
        result += int64(stack[len(stack) - 1][1]) 
    }
    return result
}