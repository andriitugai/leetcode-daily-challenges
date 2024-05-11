type worker struct {
	quality, wage int
}

// Max heap for quality
type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return x
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	workers := make([]worker, len(quality))
	for i := range workers {
		workers[i] = worker{quality[i], wage[i]}
	}

	// Sort workers by ratio of wage/quality
	sort.Slice(workers, func(i, j int) bool {
		return workers[i].wage*workers[j].quality < workers[j].wage*workers[i].quality
	})

	h, qSum, minCost := make(maxHeap, 0), 0, math.MaxFloat64
	for _, w := range workers {
		heap.Push(&h, w.quality)
		qSum += w.quality
		if h.Len() > k {
			qSum -= heap.Pop(&h).(int)
		}
		cost := float64(w.wage*qSum) / float64(w.quality)
		if h.Len() == k && cost < minCost {
			minCost = cost
		}
	}

	return minCost
}