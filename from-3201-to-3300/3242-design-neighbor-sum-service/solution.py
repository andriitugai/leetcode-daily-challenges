class neighborSum:

    def __init__(self, grid: List[List[int]]):
        self.adjacent = dict()
        self.diagonal = dict()
        n = len(grid)
        for r in range(n):
            for c in range(n):
                adj = (grid[r - 1][c] if r > 0 else 0) + \
                    (grid[r + 1][c] if r < n - 1 else 0) + \
                    (grid[r][c - 1] if c > 0 else 0) + \
                    (grid[r][c + 1] if c < n - 1 else 0)
                self.adjacent[grid[r][c]] = adj

                diag = (grid[r - 1][c - 1] if r > 0 and c > 0 else 0) + \
                    (grid[r - 1][c + 1] if r > 0 and c < n - 1 else 0) + \
                    (grid[r + 1][c - 1] if r < n - 1 and c > 0 else 0) + \
                    (grid[r + 1][c + 1] if r < n - 1 and c < n - 1 else 0)
                self.diagonal[grid[r][c]] = diag


    def adjacentSum(self, value: int) -> int:
        return self.adjacent[value]
        

    def diagonalSum(self, value: int) -> int:
        return self.diagonal[value]        


# Your neighborSum object will be instantiated and called as such:
# obj = neighborSum(grid)
# param_1 = obj.adjacentSum(value)
# param_2 = obj.diagonalSum(value)