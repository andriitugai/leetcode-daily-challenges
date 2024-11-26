func findChampion(n int, edges [][]int) int {
	incoming := make([]int, n)
	for _, edge := range edges {
		incoming[edge[1]] += 1
	}

	champion := -1
	for i := 0; i < n; i++ {
		if incoming[i] == 0 {
			if champion > -1 {
				return -1
			}
			champion = i
		}
	}
	return champion
}