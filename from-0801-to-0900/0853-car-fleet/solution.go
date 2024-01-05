func carFleet(target int, position []int, speed []int) int {
    pairs := make([][2]int, len(position))
	for i := range pairs {
		pairs[i][0] = position[i]
		pairs[i][1] = speed[i]
	}

    sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

    stack := make([]float64, 0, len(position))
    for i := range pairs {
		currentTime := float64(target-pairs[i][0]) / float64(pairs[i][1])
		// If it goes faster than the car with the tail, it means they are teaming up
		// and there is no need to fluff it in the text
		if len(stack) == 0 || currentTime > stack[len(stack)-1] {
			stack = append(stack, currentTime)
		}
	}
	return len(stack)
}