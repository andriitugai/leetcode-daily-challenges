class Solution:
    def countSquares(self, matrix: List[List[int]]) -> int:
        cache = dict()
        m, n = len(matrix), len(matrix[0])

        def dfs(r: int, c: int) -> int:
            if r == m or c == n or not matrix[r][c]:
                return 0
            if (r, c) in cache:
                return cache[(r,c)]

            cache[(r, c)] = 1 + min(dfs(r, c + 1), dfs(r + 1, c), dfs(r + 1, c + 1))
            return cache[(r, c)]

        result = 0
        for r in range(m):
            for c in range(n):
                result += dfs(r, c)
        return result
