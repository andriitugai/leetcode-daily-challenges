class Solution:
    def highestPeak(self, isWater: List[List[int]]) -> List[List[int]]:
        dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)]
        m, n = len(isWater), len(isWater[0])
        height = [[-1] * n for _ in range(m)]

        q = deque()
        for row in range(m):
            for col in range(n):
                if isWater[row][col] == 1:
                    height[row][col] = 0
                    q.append((row, col))

        while q:
            row, col = q.popleft()
            for dr, dc in dirs:
                row_, col_ = row + dr, col + dc
                if 0 <= row_ < m and 0 <= col_ < n and height[row_][col_] == -1:
                    height[row_][col_] = height[row][col] + 1
                    q.append((row_, col_))

        return height