class Solution:
    def minFallingPathSum(self, matrix: List[List[int]]) -> int:
        # Prepare DP matrix - just a copy of the matrix
        dp = [row.copy() for row in matrix]
        
        n = len(dp[0]) # the matrix is square
        for row in range(1, n):
            for col in range(n):
                # Add to the every cell in the current row
                # the min value from the 3 upper cells
                v = dp[row-1][col]
                if col > 0:
                    v = min(v, dp[row-1][col-1])
                if col < n-1:
                    v = min(v, dp[row-1][col+1])
                    
                dp[row][col] += v
        
        # The answer is the min value in the last row
        return min(dp[-1])
    