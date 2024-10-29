class Solution:
    def maxMoves(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dirs = [(0, 1), (1, 1), (-1, 1)]
        cache = dict()

        def dfs(r: int, c: int) -> int:
            if (r, c) in cache:
                return cache[(r, c)]
            result = 0
            for dr, dc in dirs:
                new_r, new_c = r + dr, c + dc
                if 0 <= new_r < m and new_c < n and grid[r][c] < grid[new_r][new_c]:
                    result = max(result, 1 + dfs(new_r, new_c))

            cache[(r, c)] = result
            return result

        return max(dfs(row, 0) for row in range(m))