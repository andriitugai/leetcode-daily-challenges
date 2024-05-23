func beautifulSubsets(nums []int, k int) int {
    count := 0
	lenNums := len(nums)
	visited := make(map[int]int)

	var explore func(index int)
	explore = func(index int) {
		if index == lenNums {
			count++
			return
		}

		num := nums[index]

		if _, exists1 := visited[num-k]; !exists1 {
			if _, exists2 := visited[num+k]; !exists2 {
				visited[num]++
				explore(index + 1)
				visited[num]--
				if visited[num] == 0 {
					delete(visited, num)
				}
			}
		}

		explore(index + 1)
	}

	explore(0)
	return count - 1 // Subtract 1 to exclude the empty subset
}