class Solution:
    def countServers(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        rows, cols = [0] * m, [0] * n
        
        servers = 0
        for r in range(m):
            for c in range(n):
                if grid[r][c]:
                    servers += 1
                    rows[r] += 1
                    cols[c] += 1

        for r in range(m):
            for c in range(n):
                if grid[r][c] and rows[r] == 1 and cols[c] == 1:
                    servers -= 1

        return servers