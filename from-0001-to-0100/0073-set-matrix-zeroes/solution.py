class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        m = len(matrix)
        n = len(matrix[0])
        
        rows = [False] * m
        cols = [False] * n
        
        for r in range(m):
            for c in range(n):
                if matrix[r][c] == 0:
                    rows[r] = True
                    cols[c] = True
                    continue
                    
        for r in range(m):
            for c in range(n):
                if rows[r] or cols[c]:
                    matrix[r][c] = 0
                    
        return matrix
                