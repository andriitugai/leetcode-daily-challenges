class Solution:
    def zigzagTraversal(self, grid: List[List[int]]) -> List[int]:
        result = []
        skipping = False
        m, n = len(grid), len(grid[0])
        for row in range(m):
            if row % 2 == 0:
                for col in range(n):
                    if not skipping:
                        result.append(grid[row][col])
                    skipping = not skipping
            else:
                for col in range(n-1, -1, -1):
                    if not skipping:
                        result.append(grid[row][col])
                    skipping = not skipping
        
        return result
