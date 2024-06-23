class Solution:
    def minimumArea(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])

        foundOne = False
        for r in range(m):
            for c in range(n):
                if grid[r][c] == 1:
                    firstRow = r
                    foundOne = True
                    break
            if foundOne:
                break

        foundOne = False
        for r in range(m - 1, -1, -1):
            for c in range(n):
                if grid[r][c] == 1:
                    lastRow = r
                    foundOne = True
                    break
            if foundOne:
                break

        foundOne = False
        for c in range(n):
            for r in range(firstRow, lastRow + 1):
                if grid[r][c] == 1:
                    firstCol = c
                    foundOne = True
            if foundOne:
                break

        foundOne = False
        for c in range(n - 1, -1, -1):
            for r in range(firstRow, lastRow + 1):
                if grid[r][c] == 1:
                    lastCol = c
                    foundOne = True
            if foundOne:
                break

        return (lastRow - firstRow + 1) * (lastCol - firstCol + 1)