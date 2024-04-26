class Solution:
    def minFallingPathSum(self, matrix: List[List[int]]) -> int:
        # Prepare DP matrix - just a copy of the matrix
        dp = [row.copy() for row in matrix]
        
        def find_two_smallest(arr: List[int]) -> (int, int, int):
            min1, min2 = float('inf'), float('inf')
            idx = -1
            for i, val in enumerate(arr):
                if val <= min1:
                    min2 = min1
                    min1 = val
                    idx = i
                elif val < min2:
                    min2 = val
            return min1, min2, idx
        
        n = len(dp[0]) # the matrix is square
        for row in range(1, n):
            m1, m2, idx = find_two_smallest(dp[row-1])
            for col in range(n):
                dp[row][col] += (m2 if col == idx else m1)
        
        # The answer is the min value in the last row
        return min(dp[-1])
    