func maximumSafenessFactor(grid [][]int) int {
    distToThief := getDistToThief(grid)
	n := len(grid)
	l, r := 0, n*2

	for l < r {
		m := (l + r) / 2
		if hasValidPath(distToThief, m) {
			l = m + 1
		} else {
			r = m
		}
	}

	return l - 1
}

func hasValidPath(distToThief [][]int, safeness int) bool {
	n := len(distToThief)
	q := make([][2]int, 0)
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	q = append(q, [2]int{0, 0})
	seen[0][0] = true

	dirs := []int{0, 1, 0, -1, 0}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		i, j := curr[0], curr[1]

		if distToThief[i][j] < safeness {
			continue
		}

		if i == n-1 && j == n-1 {
			return true
		}

		for k := 0; k < 4; k++ {
			x, y := i+dirs[k], j+dirs[k+1]

			if x >= 0 && x < n && y >= 0 && y < n && !seen[x][y] {
				q = append(q, [2]int{x, y})
				seen[x][y] = true
			}
		}
	}

	return false
}

func getDistToThief(grid [][]int) [][]int {
	n := len(grid)
	distToThief := make([][]int, n)
	q := make([][2]int, 0)
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, n)
		distToThief[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, [2]int{i, j})
				seen[i][j] = true
			}
		}
	}

	dirs := []int{0, 1, 0, -1, 0}

	for dist := 0; len(q) > 0; dist++ {
		for sz := len(q); sz > 0; sz-- {
			curr := q[0]
			q = q[1:]
			i, j := curr[0], curr[1]
			distToThief[i][j] = dist

			for k := 0; k < 4; k++ {
				x, y := i+dirs[k], j+dirs[k+1]

				if x >= 0 && x < n && y >= 0 && y < n && !seen[x][y] {
					q = append(q, [2]int{x, y})
					seen[x][y] = true
				}
			}
		}
	}

	return distToThief
}