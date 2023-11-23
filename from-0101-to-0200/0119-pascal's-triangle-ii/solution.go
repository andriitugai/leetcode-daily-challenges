func getRow(rowIndex int) []int {
	prev := []int{1}
	if rowIndex == 0 {
		return prev
	}
	var curr []int
	for i := 1; i <= rowIndex; i++ {
		curr = make([]int, i+1)
		curr[0] = 1
		curr[i] = 1
		for j := 1; j < i; j++ {
			curr[j] = prev[j-1] + prev[j]
		}
		prev = curr
	}
	return curr
}