func pathInZigZagTree(label int) []int {
    result := []int{label}
    for label != 1 {
        level := int(math.Log2(float64(label)))
		start := int(math.Pow(2, float64(level)))
		end := int(math.Pow(2, float64(level+1))) - 1
		complement := start + end - label
		label = complement / 2
		result = append([]int{label}, result...)
    }
    return result
}
