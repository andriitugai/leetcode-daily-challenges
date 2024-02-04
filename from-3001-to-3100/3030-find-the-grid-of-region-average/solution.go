func resultGrid(image [][]int, threshold int) [][]int {
    m, n := len(image), len(image[0])
    result := make([][]int, m)
    cnt := make([][]int, m)
    for i := 0; i < m; i ++ {
        result[i] = make([]int, n)
        cnt[i] = make([]int, n)
    }

    abs := func(a int) int {
        if a < 0 {
            return -a
        }
        return a
    }
    
    isValid := func(x, y int) bool {
        for _, dx := range []int{-1, 1} {
            for _, dy := range []int{-1, 1} {
                if abs(image[x][y] - image[x+dx][y]) > threshold || abs(image[x][y] - image[x][y+dy]) > threshold {
                    return false
                }
            }
        }
        if abs(image[x-1][y-1] - image[x-1][y]) > threshold || abs(image[x-1][y-1] - image[x][y-1]) > threshold {
            return false
        }
        if abs(image[x-1][y+1] - image[x-1][y]) > threshold || abs(image[x-1][y+1] - image[x][y+1]) > threshold {
            return false
        }
        if abs(image[x+1][y-1] - image[x+1][y]) > threshold || abs(image[x+1][y-1] - image[x][y-1]) > threshold {
            return false
        }
        if abs(image[x+1][y+1] - image[x+1][y]) > threshold || abs(image[x+1][y+1] - image[x][y+1]) > threshold {
            return false
        }
        return true
    }

    for i := 1; i < m - 1; i ++ {
        for j := 1; j < n - 1; j ++ {
            if isValid(i, j) {
                val := 0
                for _, dx := range []int{-1, 0, 1} {
                    for _, dy := range []int{-1, 0, 1} {
                        cnt[i+dx][j+dy] += 1
                        val += image[i+dx][j+dy]
                    }
                }
                for _, dx := range []int{-1, 0, 1} {
                    for _, dy := range []int{-1, 0, 1} {
                        result[i+dx][j+dy] += (val / 9)
                    }
                }
            }
        }
    }

    for i := 0; i < m; i ++ {
        for j := 0; j < n; j ++ {
            if cnt[i][j] == 0 {
                result[i][j] = image[i][j]
            } else {
                result[i][j] /= cnt[i][j]
            }
        }
    }
    return result
}