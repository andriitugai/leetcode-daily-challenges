class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        heap = []
        for num in nums:
            if num < k:
                heapq.heappush(heap, num)

        result = 0
        while heap:
            result += 1
            a = heapq.heappop(heap)
            if heap:
                b = heapq.heappop(heap)
                c = 2 * a + b
                if c < k:
                    heapq.heappush(heap, c)

        return result