class Solution:
    def sortMatrix(self, grid: List[List[int]]) -> List[List[int]]:
        n = len(grid)
        diags = defaultdict(list)
        for row in range(n):
            for col in range(n):
                diags[row - col].append(grid[row][col])

        for idx, diag in diags.items():
            if idx < 0:
                diags[idx] = sorted(diag, reverse=False)
            else:
                diags[idx] = sorted(diag, reverse=True)

        for row in range(n):
            for col in range(n):
                grid[row][col] = diags[row - col].pop(0)

        return grid
        