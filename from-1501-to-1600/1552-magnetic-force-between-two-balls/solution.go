func maxDistance(position []int, m int) int {
    n := len(position)
    sort.Ints(position)

    countBalls := func(dist int) int {
        numBalls := 1
        curr := position[0]
        for _, pos := range position[1:] {
            if pos - curr >= dist {
                numBalls += 1
                curr = pos
            }
        }
        return numBalls
    }

    left, right := 1, position[n - 1] - position[0]
    for left <= right {
        mid := left + (right - left) / 2
        if countBalls(mid) >= m {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return right
}