func imageSmoother(img [][]int) [][]int {
    m, n  := len(img), len(img[0])

    smoothed := make([][]int, m)
    
    for r := 0; r < m; r++ {
        smoothed[r] = make([]int, n)
        for c := 0; c < n; c++ {
            cnt := 0
            sum := 0
            for i := r - 1; i < r + 2; i++ {
                if i >= 0 && i < m {
                    for j := c - 1; j < c + 2; j++ {
                        if j >= 0 && j < n {
                            sum += img[i][j]
                            cnt += 1
                        }
                    }
                }
            }
            smoothed[r][c] = sum / cnt
        }
    }
    return smoothed
}