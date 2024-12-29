class Solution:
    def minimumOperations(self, grid: List[List[int]]) -> int:
        result = 0
        for col in range(len(grid[0])):
            prev = grid[0][col]
            for row in range(1, len(grid)):
                if prev >= grid[row][col]:
                    result += prev - grid[row][col] + 1
                    grid[row][col] = prev + 1
                prev = grid[row][col]
        return result