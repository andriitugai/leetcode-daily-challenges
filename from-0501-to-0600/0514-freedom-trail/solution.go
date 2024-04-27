func findRotateSteps(ring string, key string) int {
	var m, n = len(ring), len(key)
	var pos = make(map[uint8][]int)
	for i := 0; i < len(ring); i++ {
		pos[ring[i]] = append(pos[ring[i]], i)
	}
	var mem = make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i, j int) int // ring[i] is currently at 12 o'clock , j stands for current key[j]
	dfs = func(i, j int) (res int) {
		if j == n {
			return
		}
		var p = &mem[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		res = math.MaxInt
		for _, po := range pos[key[j]] { // scan all the position key[j] in ring
			res = min(res, dfs(po, j+1)+min(abs(i-po), m-abs(i-po))) // we need to make po at 12 o'clock by rotating ring clockwise or counterclockwise (choose the smaller one), and then plus dfs(po,j+1)
		}
		return
	}
	return n + dfs(0, 0) // every char in key needs to be pressed too
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}