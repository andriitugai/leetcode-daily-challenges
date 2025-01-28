class Solution:
    def findMaxFish(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)]
        visited = set()

        def dfs(x, y):
            if x < 0 or x >= m or y < 0 or y >= n or grid[x][y] == 0 or (x, y) in visited:
                return 0

            visited.add((x, y))
            tot = grid[x][y]
            for dx, dy in dirs:
                tot += dfs(x + dx, y + dy)

            return tot

        result = 0
        for row in range(m):
            for col in range(n):
                if grid[row][col] > 0 and (row, col) not in visited:
                    result = max(result, dfs(row, col))

        return result
