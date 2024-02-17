import heapq

class Solution:
    def furthestBuilding(self, heights: List[int], bricks: int, ladders: int) -> int:
        heap = []

        bricks_count = 0

        for i in range(len(heights) - 1):
            diff = heights[i+1] - heights[i]
            if diff > 0:
                bricks_count += diff
                heapq.heappush(heap, -diff)

                if bricks_count > bricks:
                    if ladders > 0:
                        ladders -= 1
                        bricks_count += heapq.heappop(heap)
                    else:
                        return i
                    
        return len(heights) - 1