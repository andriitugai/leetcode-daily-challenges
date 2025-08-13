class Solution:
    def reverseSubmatrix(self, grid: List[List[int]], x: int, y: int, k: int) -> List[List[int]]:
        for col in range(y, y + k):
            top, bottom = x, x + k - 1
            while top < bottom:
                grid[top][col], grid[bottom][col] = grid[bottom][col], grid[top][col]
                top += 1
                bottom -= 1
        return grid