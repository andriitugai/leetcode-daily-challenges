class Solution:
    def getFinalState(self, nums: List[int], k: int, multiplier: int) -> List[int]:
        result = nums[::]
        min_heap = [(num, i) for i, num in enumerate(nums)]
        heapq.heapify(min_heap)

        for _ in range(k):
            n, i = heapq.heappop(min_heap)
            result[i] *= multiplier

            heapq.heappush(min_heap, (result[i], i))

        return result