class Solution:
    def pickGifts(self, gifts: List[int], k: int) -> int:
        gifts = [-g for g in gifts]
        heapq.heapify(gifts)

        for _ in range(k):
            g = -1 * math.floor(math.sqrt(-1 * heapq.heappop(gifts)))
            heapq.heappush(gifts, g)

        return -sum(gifts)