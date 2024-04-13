func maximalRectangle(matrix [][]byte) int {
    m, n := len(matrix), len(matrix[0])
    result := 0
    height := make([]int, n + 1)
    for r := 0; r < m; r ++ {
        for c := 0; c < n; c ++ {
            if matrix[r][c] == '1' {
                height[c] += 1
            } else {
                height[c] = 0
            }
        }
        stack := []int{}
        for c := 0; c <= n; c ++ {
            for len(stack) > 0 && height[c] < height[stack[len(stack) - 1]] {
                h := height[stack[len(stack) - 1]]
                stack = stack[:len(stack) - 1]
                w := c
                if len(stack) > 0 {
                    w = c - stack[len(stack) - 1] - 1
                }
                result = max(result, h * w)
            }
            stack = append(stack, c)
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}