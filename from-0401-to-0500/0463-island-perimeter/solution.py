class Solution:
    def islandPerimeter(self, grid: List[List[int]]) -> int:
        m = len(grid)
        n = len(grid[0])
        
        visited = set()
        
        def dfs(row, col):
            nonlocal grid, m, n, visited
            if row < 0 or col < 0 or row >= m or col >= n or grid[row][col] == 0:
                return 1
            if (row, col) in visited:
                return 0
            visited.add((row, col))

            perim = dfs(row + 1, col)
            perim += dfs(row - 1, col)
            perim += dfs(row, col + 1)
            perim += dfs(row, col - 1)
            
            return perim
        
        for row in range(m):
            for col in range(n):
                if grid[row][col] == 1:
                    return dfs(row, col)
                
        return 0
            