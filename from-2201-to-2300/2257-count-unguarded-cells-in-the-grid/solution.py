class Solution:
    def countUnguarded(self, m: int, n: int, guards: List[List[int]], walls: List[List[int]]) -> int:
        grid = [[0] * n for _ in range(m)]
        # 0 - free, 1 - guard, 2 - wall, 3 - guardable

        for r, c in guards:
            grid[r][c] = 1
        for r, c in walls:
            grid[r][c] = 2

        for r, c in guards:
            for i in range(r - 1, -1, -1):
                if grid[i][c] == 2 or grid[i][c] == 1:
                    break
                grid[i][c] = 3
            for i in range(r + 1, m, 1):
                if grid[i][c] == 2 or grid[i][c] == 1:
                    break
                grid[i][c] = 3
            for j in range(c - 1, -1, -1):
                if grid[r][j] == 2 or grid[r][j] == 1:
                    break
                grid[r][j] = 3
            for j in range(c + 1, n, 1):
                if grid[r][j] == 2 or grid[r][j] == 1:
                    break
                grid[r][j] = 3

        result = 0
        for r in range(m):
            for c in range(n):
                if grid[r][c] == 0:
                    result += 1

        return result