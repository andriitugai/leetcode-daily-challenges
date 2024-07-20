class Solution:
    def restoreMatrix(self, rowSum: List[int], colSum: List[int]) -> List[List[int]]:
        m, n  = len(rowSum), len(colSum)
        ans = [[0] * n for _ in range(m)]
        
        for r in range(m):
            for c in range(n):
                ans[r][c] = min(rowSum[r], colSum[c])
                rowSum[r] -= ans[r][c]
                colSum[c] -= ans[r][c]
                
        return ans