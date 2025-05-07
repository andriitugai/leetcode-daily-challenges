class Solution:
    def minTimeToReach(self, moveTime: List[List[int]]) -> int:
        m, n = len(moveTime), len(moveTime[0])
        visited = set()

        dirs = [(1, 0), (-1, 0), (0, 1), (0, -1)]
        
        pq = [(0, 0, 0)]
        heapify(pq)

        while pq:
            time, row, col = heappop(pq)
            if row == m - 1 and col == n - 1:
                return time

            for dr, dc in dirs:
                row_, col_ = row + dr, col + dc
                if row_ < 0 or row_ >= m or col_ < 0 or col_ >= n or (row_, col_) in visited:
                    continue
                visited.add((row_, col_))

                time_ = max(time, moveTime[row_][col_]) + 1
                heappush(pq, (time_, row_, col_))

        return -1
