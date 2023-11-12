func numBusesToDestination(routes [][]int, source int, target int) int {
    if source == target {
        return 0
    }

    stopToBuses := make(map[int][]int)
    for bus, stops := range routes {
        for _, stop := range stops {
            stopToBuses[stop] = append(stopToBuses[stop], bus)
        }
    }

    q := []int{}
    visitedBuses := make(map[int]bool)
    visitedStops := make(map[int]bool)

    q = append(q, source)
    visitedStops[source] = true
    busesTaken := 0

    for len(q) > 0 {
        busesTaken += 1
        n := len(q)
        for i := 0; i < n; i ++ {
            currStop := q[0]
            q = q[1:]

            for _, bus := range stopToBuses[currStop] {
                if visitedBuses[bus] {
                    continue
                }
                visitedBuses[bus] = true
                for _, stop := range routes[bus] {
                    if stop == target {
                        return busesTaken
                    }
                    if visitedStops[stop] {
                        continue
                    }
                    visitedStops[stop] = true

                    q = append(q, stop)
                }
            }
        }
    }
    return -1
}