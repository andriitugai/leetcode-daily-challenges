func findContentChildren(g []int, s []int) int {
    sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
    sort.Slice(g,func(i,j int) bool{
        return g[i] > g[j]
    })

    result, j := 0, 0
    for i := 0; i < len(g) && j < len(s); i += 1 {
        if g[i] <= s[j] {
            result += 1
            j += 1
        }
    }
    return result
}